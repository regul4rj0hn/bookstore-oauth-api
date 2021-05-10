package token

import (
	"net/http"
	"strings"
	"time"

	"github.com/regul4rj0hn/bookstore-oauth-api/pkg/domain/errors"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"exp"`
	Subject     int64  `json:"sub"`
}

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at AccessToken) Validate() *errors.Response {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.BadRequest("invalid access token")
	}
	if at.Subject <= 0 {
		return errors.BadRequest("invalid sub")
	}
	if at.ClientId <= 0 {
		return errors.BadRequest("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.BadRequest("invalid expiration date")
	}
	return nil
}

func (*AccessToken) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
