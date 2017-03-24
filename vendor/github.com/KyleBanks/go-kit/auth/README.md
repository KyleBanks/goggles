# auth
--
    import "github.com/KyleBanks/go-kit/auth/"

Package auth provides generic authentication functionality.

Note this is not 100% secure and should only be used for prototyping, not for
production systems or systems that are accessed by real users.

## Usage

#### func  GetIdentifierForAccessToken

```go
func GetIdentifierForAccessToken(a string) (string, error)
```
GetIdentifierForAccessToken returns a user's identifier, as returned by the
Authenticator interface, if it exists in the cache.

If the identifier does not exist, and empty string and error will be returned.

#### func  GetIdentifierForRefreshToken

```go
func GetIdentifierForRefreshToken(r string) (string, error)
```
GetIdentifierForRefreshToken returns a user's identifier, as returned by the
Authenticator interface, if it exists in the cahce.

If the identifier does not exist, an empty string and error will be returned.

#### func  HashPassword

```go
func HashPassword(plainText string) (string, error)
```
HashPassword returns a hashed version of the plain-text password provided.

#### func  SetCache

```go
func SetCache(c cache.Cacher)
```
SetCache sets the Cache to use for authentication tokens.

#### type AuthenticationTokenPair

```go
type AuthenticationTokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
```

AuthenticationTokenPair represents a pair of authentication tokens (Access and
Refresh).

#### func  Authenticate

```go
func Authenticate(a Authenticator, plainTextPassword string) (AuthenticationTokenPair, error)
```
Authenticate validates an Authenticator based on it's password hash and the
plain-text password provided.

#### func  GenerateToken

```go
func GenerateToken() AuthenticationTokenPair
```
GenerateToken returns a new AuthenticationTokenPair

#### func  Refresh

```go
func Refresh(a Authenticator, refreshToken string) (AuthenticationTokenPair, error)
```
Refresh generates a new token pair for a given authenticator.

#### type Authenticator

```go
type Authenticator interface {
	Identifier() string     // Identifier returns a unique reference to this user.
	HashedPassword() string // HashedPassword returns the user's password hash.
}
```

Authenticator defines an interface for an authenticate-able User.
