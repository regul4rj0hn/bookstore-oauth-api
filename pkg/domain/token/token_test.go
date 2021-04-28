package access_token

import (
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("default token expiration time should be 24 hours")
	}
}

func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()
	if at.IsExpired() {
		t.Error("new access token should not be expired")
	}

	if at.Token != "" {
		t.Error("new access token should not have valid id")
	}

	if at.Subject != 0 {
		t.Error("new access token should not have a subject associated")
	}
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token expiring three hours from now should not be expired")
	}
}
