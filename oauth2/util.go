package oauth2

import (
	"context"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
)

func GetAccessTokenFromContext(ctx context.Context) (string, error) {
	ts, ok := ctx.Value("tokenSource").(oauth2.TokenSource)
	if !ok {
		return "", errors.New("tokenSource is not of type oauth2.TokenSource")
	}

	token, err := ts.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map,
// since we'll only be modify the headers.
// Per the specification of http.RoundTripper, we should not directly modify a request.
func cloneRequest(r *http.Request) *http.Request {
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

// Sets the User-Agent header for requests.
// We need to set a custom user agent because using the one set by the
// stdlib gives us 429 Too Many Requests responses from the Reddit API.
type userAgentTransport struct {
	userAgent string
	Base      http.RoundTripper
}

func (t *userAgentTransport) setUserAgent(req *http.Request) *http.Request {
	req2 := cloneRequest(req)
	req2.Header.Set("user-agent", t.userAgent)
	return req2
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := t.setUserAgent(req)
	return t.base().RoundTrip(req2)
}

func (t *userAgentTransport) base() http.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return http.DefaultTransport
}
