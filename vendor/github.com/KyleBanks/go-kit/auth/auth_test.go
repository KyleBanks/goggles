package auth

import (
	"testing"

	"github.com/KyleBanks/go-kit/cache"
)

//--------------------------------------------------
// For testing purposes only!
//--------------------------------------------------
type testAuthenticator struct {
	ID   string
	Hash string
}

func (t testAuthenticator) Identifier() string {
	return t.ID
}
func (t testAuthenticator) HashedPassword() string {
	return t.Hash
}

//--------------------------------------------------

func TestSetCache(t *testing.T) {
	SetCache(cache.New("localhost:6379"))
}

func TestHashPassword(t *testing.T) {
	testPassword := "flexcry69"

	// Basic use case
	hash, err := HashPassword(testPassword)
	if err != nil {
		t.Error(err)
	} else if len(hash) == 0 {
		t.Error("Hash should not be 0 length!")
	}

	// A different password should return a different hash
	if different, err := HashPassword("different"); err != nil {
		t.Error(err)
	} else if different == hash {
		t.Error("Hashes must be different for different plain-text passwords:", different, hash)
	}
}

func TestAuthenticate(t *testing.T) {
	// Hash a password and instantiate a fake "user"
	hash, err := HashPassword("flexcry69")
	if err != nil {
		t.Error(err)
	}
	authenticator := testAuthenticator{
		ID:   "134",
		Hash: hash,
	}

	// Authenticate the user with the valid password
	if tokens, err := Authenticate(authenticator, "flexcry69"); err != nil {
		t.Error(err)
	} else if tokens == (AuthenticationTokenPair{}) {
		t.Error("Empty token pair returned!")
	}

	// Authenticate with an invalid password
	if _, err := Authenticate(authenticator, "badpass"); err == nil {
		t.Error("Expected error authenticating with bad password!")
	}
}

func TestRefresh_HappyPath(t *testing.T) {
	// Hash a password and instantiate a fake "user"
	hash, err := HashPassword("flexcry69")
	if err != nil {
		t.Error(err)
	}
	authenticator := testAuthenticator{
		ID:   "134",
		Hash: hash,
	}

	// Authenticate the user with the valid password
	tokens, err := Authenticate(authenticator, "flexcry69")
	if err != nil {
		t.Error(err)
	}

	// Refresh the tokens
	newTokens, err := Refresh(authenticator, tokens.RefreshToken)
	if err != nil {
		t.Error(err)
	} else if newTokens == (AuthenticationTokenPair{}) {
		t.Error("Unexpected empty token pair!")
	}

	// Validate them
	if newTokens.AccessToken == tokens.AccessToken || newTokens.RefreshToken == tokens.RefreshToken {
		t.Error("Refreshed tokens should have different values than the existing ones:", newTokens, tokens)
	}
}

func TestRefresh_InvalidTokens(t *testing.T) {
	// Hash a password and instantiate a fake "user"
	hash, err := HashPassword("flexcry69")
	if err != nil {
		t.Error(err)
	}
	authenticator := testAuthenticator{
		ID:   "134",
		Hash: hash,
	}

	// Attempt to refresh with a fake token
	if _, err := Refresh(authenticator, "fake token"); err == nil {
		t.Error("Expected error refreshing fake tokens!")
	}
}

func TestGetIdentifierForAccessToken_HappyPath(t *testing.T) {
	// Hash a password and instantiate a fake "user"
	hash, err := HashPassword("flexcry69")
	if err != nil {
		t.Error(err)
	}
	authenticator := testAuthenticator{
		ID:   "134",
		Hash: hash,
	}

	// Authenticate the user with the valid password
	tokens, err := Authenticate(authenticator, "flexcry69")
	if err != nil {
		t.Error(err)
	}

	// Get their identifier and compare to the authenticator we created
	if identifier, err := GetIdentifierForAccessToken(tokens.AccessToken); err != nil {
		t.Error(err)
	} else if identifier != authenticator.Identifier() {
		t.Error("Unexpected identifier returned: ", identifier, "Expected:", authenticator.Identifier())
	}
}

func TestGetIdentifierForAccessToken_BadToken(t *testing.T) {
	if _, err := GetIdentifierForAccessToken("fake token"); err == nil {
		t.Error("Expected error for invalid access token!")
	}
}

func TestGetIdentifierForRefreshToken_HappyPath(t *testing.T) {
	// Hash a password and instantiate a fake "user"
	hash, err := HashPassword("flexcry69")
	if err != nil {
		t.Error(err)
	}
	authenticator := testAuthenticator{
		ID:   "134",
		Hash: hash,
	}

	// Authenticate the user with the valid password
	tokens, err := Authenticate(authenticator, "flexcry69")
	if err != nil {
		t.Error(err)
	}

	// Get their identifier and compare to the authenticator we created
	if identifier, err := GetIdentifierForRefreshToken(tokens.RefreshToken); err != nil {
		t.Error(err)
	} else if identifier != authenticator.Identifier() {
		t.Error("Unexpected identifier returned: ", identifier, "Expected:", authenticator.Identifier())
	}
}

func TestGetIdentifierForRefreshToken_BadToken(t *testing.T) {
	if _, err := GetIdentifierForRefreshToken("fake token"); err == nil {
		t.Error("Expected error for invalid refresh token!")
	}
}
