package auth

import (
	"github.com/satori/go.uuid"
)

// AuthenticationTokenPair represents a pair of authentication tokens (Access and Refresh).
type AuthenticationTokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// GenerateToken returns a new AuthenticationTokenPair
func GenerateToken() AuthenticationTokenPair {
	return AuthenticationTokenPair{
		AccessToken:  uuid.NewV4().String(),
		RefreshToken: uuid.NewV4().String(),
	}
}
