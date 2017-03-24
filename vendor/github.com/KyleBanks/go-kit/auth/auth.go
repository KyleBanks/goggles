// Package auth provides generic authentication functionality.
//
// Note this is not 100% secure and should only be used for prototyping,
// not for production systems or systems that are accessed by real users.
package auth

import (
	"fmt"

	"github.com/KyleBanks/go-kit/cache"
	"golang.org/x/crypto/bcrypt"
)

var (
	authCache cache.Cacher // cache is used for storing authentication tokens.
)

// Authenticator defines an interface for an authenticate-able User.
type Authenticator interface {
	Identifier() string     // Identifier returns a unique reference to this user.
	HashedPassword() string // HashedPassword returns the user's password hash.
}

// SetCache sets the Cache to use for authentication tokens.
func SetCache(c cache.Cacher) {
	authCache = c
}

// HashPassword returns a hashed version of the plain-text password provided.
func HashPassword(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash[:]), nil
}

// Authenticate validates an Authenticator based on it's password hash and the plain-text
// password provided.
func Authenticate(a Authenticator, plainTextPassword string) (AuthenticationTokenPair, error) {
	err := bcrypt.CompareHashAndPassword([]byte(a.HashedPassword()), []byte(plainTextPassword))
	if err != nil {
		return AuthenticationTokenPair{}, err
	}

	// Generate and cache a new token pair for this session
	return generateAndStoreTokens(a)
}

// Refresh generates a new token pair for a given authenticator.
func Refresh(a Authenticator, refreshToken string) (AuthenticationTokenPair, error) {
	newTokens, err := generateAndStoreTokens(a)
	if err != nil {
		return AuthenticationTokenPair{}, err
	}

	// Clear the old tokens from the cache
	if err := clearCachedTokens(refreshToken); err != nil {
		return AuthenticationTokenPair{}, err
	}

	return newTokens, nil
}

// GetIdentifierForAccessToken returns a user's identifier, as returned by
// the Authenticator interface, if it exists in the cache.
//
// If the identifier does not exist, and empty string and error will be returned.
func GetIdentifierForAccessToken(a string) (string, error) {
	return authCache.GetString(getAccessTokenCacheKey(a))
}

// GetIdentifierForRefreshToken returns a user's identifier, as returned by
// the Authenticator interface, if it exists in the cahce.
//
// If the identifier does not exist, an empty string and error will be returned.
func GetIdentifierForRefreshToken(r string) (string, error) {
	return authCache.GetString(getRefreshTokenCacheKey(r))
}

// generateAndStoreTokens creates and caches a new AuthenticationTokenPair.
func generateAndStoreTokens(a Authenticator) (AuthenticationTokenPair, error) {
	t := GenerateToken()
	if err := cacheTokens(t, a); err != nil {
		return AuthenticationTokenPair{}, err
	}

	return t, nil
}

// cacheTokens stores an access token and refresh token pair for an authenticated User.
func cacheTokens(t AuthenticationTokenPair, a Authenticator) error {
	if _, err := authCache.PutString(getAccessTokenCacheKey(t.AccessToken), a.Identifier()); err != nil {
		return err
	}

	if _, err := authCache.PutString(getRefreshTokenCacheKey(t.RefreshToken), a.Identifier()); err != nil {
		return err
	}

	if _, err := authCache.PutString(getRefreshToAccessTokenCacheKey(t.RefreshToken), t.AccessToken); err != nil {
		return err
	}

	return nil
}

// getAccessTokenCacheKey returns the access token cache key.
func getAccessTokenCacheKey(accessToken string) string {
	return fmt.Sprintf("accessToken:%s", accessToken)
}

// getRefreshTokenCacheKey returns the refresh token cache key.
func getRefreshTokenCacheKey(refreshToken string) string {
	return fmt.Sprintf("refreshToken:%s", refreshToken)
}

// getRefreshToAccessTokenCacheKey returns the refresh -> access token cache key.
func getRefreshToAccessTokenCacheKey(refreshToken string) string {
	return fmt.Sprintf("refreshToAccessToken:%s", refreshToken)
}

// clearCachedTokens clears all tokens associated to a refresh token.
func clearCachedTokens(r string) error {
	if a, err := authCache.GetString(getRefreshToAccessTokenCacheKey(r)); err != nil {
		return err
	} else if err = authCache.Delete(getAccessTokenCacheKey(a)); err != nil {
		return err
	} else if err = authCache.Delete(getRefreshTokenCacheKey(r)); err != nil {
		return err
	} else if err = authCache.Delete(getRefreshToAccessTokenCacheKey(r)); err != nil {
		return err
	}

	return nil
}
