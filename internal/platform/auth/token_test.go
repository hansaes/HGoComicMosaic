package auth

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTokenService(t *testing.T) {
	srv := NewTokenService([]byte("EAor0Z9jyBz1BjMB9URHFpUrcekZh9C1TVvZK31Habb"), "issuer", 24*time.Hour)

	token, _, err := srv.GenerateToken(1, "admin", true)
	assert.Nil(t, err)

	claims, err := srv.VerifyToken(token)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), claims.UserID)
	assert.Equal(t, "admin", claims.Username)
	assert.Equal(t, true, claims.IsAdmin)

}
