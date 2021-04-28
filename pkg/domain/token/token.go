package access_token

import "time"

const (
	expirationTime = 24
)

type AccessToken struct {
	Token    string `json:"access_token"`
	ClientId int64  `json:"client_id"`
	Expires  int64  `json:"exp"`
	Subject  int64  `json:"sub"`
}

func GetAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
