package auth

import (
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken()

	if len(token.AccessToken) == 0 || len(token.RefreshToken) == 0 {
		t.Error("Tokens should not be empty:", token)
	} else if token.AccessToken == token.RefreshToken {
		t.Error("Tokens must not match eachother:", token)
	}
}
