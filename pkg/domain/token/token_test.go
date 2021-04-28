package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "default token expiration time should be 24 hours")
}

func TestGetAccessToken(t *testing.T) {
	at := GetAccessToken()
	assert.False(t, at.IsExpired(), "new access token should not be expired")
	assert.EqualValues(t, "", at.Token, "new access token should not have valid id")
	assert.True(t, at.Subject == 0, "new access token should not have a subject associated")
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should not be expired")
}
