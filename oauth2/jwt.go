package oauth2

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Sadzeih/valcompbot/utils"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Token *oauth2.Token `json:"token"`
}

func (h *Repository) Middleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	bearer := r.Header.Get("Authorization")
	if len(bearer) == 0 {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	bearerParts := strings.Fields(bearer)
	if len(bearerParts) != 2 && strings.ToLower(bearerParts[0]) != "bearer" {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	token, err := jwt.ParseWithClaims(
		bearerParts[1],
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(h.jwtSigningKey), nil
		},
	)
	if err != nil {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		utils.WriteError(rw, utils.APIError{
			Code:    http.StatusUnauthorized,
			Message: "Unauthorized",
		})
		return
	}

	fmt.Println(claims.Token)

	// TokenSource is used to allow automatic refresh of token
	ts := h.config.TokenSource(r.Context(), claims.Token)

	err = h.checkIfModerator(r.Context(), claims.Token)
	if err != nil {
		return
	}

	// TokenSource is added to the context for easy use in handlers
	r = r.WithContext(context.WithValue(r.Context(), "tokenSource", ts))

	next(rw, r)
}

func (h *Repository) checkIfModerator(ctx context.Context, token *oauth2.Token) error {
	client := h.config.Client(ctx, token)

	req, err := http.NewRequest(http.MethodGet, "https://oauth.reddit.com/subreddits/mine/moderator", nil)
	if err != nil {
		return err
	}
	req.Header.Add("user-agent", "golang:valcompbot:v0.1.0")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
