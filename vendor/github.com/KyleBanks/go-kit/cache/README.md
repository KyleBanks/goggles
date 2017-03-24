# cache
--
    import "github.com/KyleBanks/go-kit/cache/"

Package cache is a simple cache wrapper, used to abstract Redis/Memcache/etc
behind a reusable API for simple use cases.

The idea is that Redis could be swapped for another cache and the client
wouldn't need to update another (except perhaps calls to New to provide
different connection parameters).

For now cache supports only Redis, but eventually that could be provided by the
client.

## Usage

```go
var (
	// ErrCantUnlock is returned if the cache fails to unlock a key.
	ErrCantUnlock = errors.New("failed to unlock")
)
```

#### type Cache

```go
type Cache struct {
}
```

Cache implements the Cacher interface using a Redis pool.

#### func  New

```go
func New(host string) *Cache
```
New instantiates and returns a new Cache.

#### func (Cache) Delete

```go
func (c Cache) Delete(key string) error
```
Delete removes an item from the cache by it's key.

#### func (Cache) Expire

```go
func (c Cache) Expire(key string, seconds time.Duration) error
```
Expire sets the time for a key to expire in seconds.

#### func (Cache) GetMarshaled

```go
func (c Cache) GetMarshaled(key string, v interface{}) error
```
GetMarshaled retrieves an item from the cache with the specified key, and
un-marshals it from JSON to the value provided.

If they key doesn't exist, an error is returned.

#### func (Cache) GetString

```go
func (c Cache) GetString(key string) (string, error)
```
GetString returns the string value stored with the given key.

If the key doesn't exist, an error is returned.

#### func (Cache) Lock

```go
func (c Cache) Lock(key, value string, timeoutMs int) (bool, error)
```
Lock attempts to put a lock on the key for a specified duration (in
milliseconds). If the lock was successfully acquired, true will be returned.

Note: The value provided can be anything, so long as it's unique. The value will
then be used when attempting to Unlock, and will only work if the value matches.
It's important that each instance that tries to perform a Lock have it's own
unique key so that you don't unlock another instances lock!

#### func (Cache) PutMarshaled

```go
func (c Cache) PutMarshaled(key string, value interface{}) (interface{}, error)
```
PutMarshaled stores a json marshalled value with the given key.

#### func (Cache) PutString

```go
func (c Cache) PutString(key string, value string) (interface{}, error)
```
PutString stores a simple key-value pair in the cache.

#### func (Cache) Unlock

```go
func (c Cache) Unlock(key, value string) error
```
Unlock attempts to remove the lock on a key so long as the value matches. If the
lock cannot be removed, either because the key has already expired or because
the value was incorrect, an error will be returned.

#### type Cacher

```go
type Cacher interface {
	PutString(key string, value string) (interface{}, error)
	GetString(key string) (string, error)

	PutMarshaled(key string, value interface{}) (interface{}, error)
	GetMarshaled(key string, v interface{}) error

	Delete(key string) error
	Expire(key string, seconds time.Duration) error

	Lock(key, value string, timeoutMs int) (bool, error)
	Unlock(key, value string) error
}
```

Cacher defines a mockable Cache interface that can store values in a key-value
cache.

#### type Mock

```go
type Mock struct {
}
```

Mock provides a mocked Cache implementation for testing.

#### func  NewMock

```go
func NewMock() *Mock
```
NewMock instantiates and returns a new Mock cache.

#### func (Mock) Delete

```go
func (m Mock) Delete(key string) error
```
Delete removes an item from the mock by it's key.

#### func (Mock) Expire

```go
func (m Mock) Expire(key string, seconds time.Duration) error
```
Expire sets the time for a key to expire in seconds.

#### func (Mock) GetMarshaled

```go
func (m Mock) GetMarshaled(key string, v interface{}) error
```
GetMarshaled retrieves an item from the cache with the specified key, and
un-marshals it from JSON to the value provided.

If they key doesn't exist, an error is returned.

#### func (Mock) GetString

```go
func (m Mock) GetString(key string) (string, error)
```
GetString returns the string value stored with the given key.

If the key doesn't exist, an error is returned.

#### func (Mock) Lock

```go
func (m Mock) Lock(key, value string, durationMs int) (bool, error)
```
Lock attempts to put a lock on the key for a specified duration (in
milliseconds).

#### func (Mock) PutMarshaled

```go
func (m Mock) PutMarshaled(key string, value interface{}) (interface{}, error)
```
PutMarshaled stores a json marshalled value with the given key.

#### func (Mock) PutString

```go
func (m Mock) PutString(key string, value string) (interface{}, error)
```
PutString stores a simple key-value pair in the mock.

#### func (Mock) Unlock

```go
func (m Mock) Unlock(key, value string) error
```
Unlock attempts to remove the lock on a key so long as the value matches.
