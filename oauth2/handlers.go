package oauth2

import (
	"github.com/Sadzeih/valcompbot/utils"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"golang.org/x/oauth2"
	"net/http"
)

type Repository struct {
	reddit        *reddit.Client
	config        *oauth2.Config
	jwtSigningKey string
}

func New(reddit *reddit.Client, jwtSigningKey, clientID, clientSecret string, scopes []string, redirectURL string) *Repository {
	r := &Repository{
		reddit:        reddit,
		jwtSigningKey: jwtSigningKey,
	}
	r.config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.reddit.com/api/v1/authorize",
			TokenURL: "https://www.reddit.com/api/v1/access_token",
		},
		RedirectURL: redirectURL,
	}

	return r
}

type AuthCodeURLResponse struct {
	URL string `json:"url"`
}

func (h *Repository) AuthCodeURLHandler(rw http.ResponseWriter, r *http.Request) {
	state := r.URL.Query().Get("state")
	if state == "" {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusBadRequest,
			Message: "state was not provided",
		})
		return
	}

	u := h.config.AuthCodeURL(state)

	utils.WriteJSON(rw, AuthCodeURLResponse{
		URL: u,
	})
}

func (h *Repository) Callback(rw http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusBadRequest,
			Message: "code was not provided",
		})
		return
	}

	token, err := h.config.Exchange(r.Context(), code)
	if err != nil {
		apiErr := utils.APIError{
			Code:    http.StatusInternalServerError,
			Message: "token could not be retrieved",
		}
		utils.WriteError(rw, apiErr)
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		Token: token,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(token.Expiry),
		},
	})

	ss, err := jwtToken.SignedString([]byte(h.jwtSigningKey))
	if err != nil {
		errMsg := "signing jwt failed"
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusInternalServerError,
			Message: errMsg,
		})
		return
	}
	_, err = rw.Write([]byte(ss))
	if err != nil {
		utils.WriteError(rw, utils.InternalServerError)
		return
	}
}
