// Package cache is a simple cache wrapper, used to abstract Redis/Memcache/etc behind a reusable
// API for simple use cases.
//
// The idea is that Redis could be swapped for another cache and the client wouldn't
// need to update another (except perhaps calls to New to provide different connection
// parameters).
//
// For now cache supports only Redis, but eventually that could be provided by the client.
package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	lockScript = `
		return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2])
	`
	unlockScript = `
		if redis.call("get",KEYS[1]) == ARGV[1] then
		    return redis.call("del",KEYS[1])
		else
		    return 0
		end
	`
)

var (
	// ErrCantUnlock is returned if the cache fails to unlock a key.
	ErrCantUnlock = errors.New("failed to unlock")
)

// Cacher defines a mockable Cache interface that can store values in a key-value cache.
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

// Cache implements the Cacher interface using a Redis pool.
type Cache struct {
	pool *redis.Pool
}

// New instantiates and returns a new Cache.
func New(host string) *Cache {
	return &Cache{
		pool: &redis.Pool{
			MaxIdle:   5,
			MaxActive: 100,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", host)
			},
		},
	}
}

// PutString stores a simple key-value pair in the cache.
func (c Cache) PutString(key string, value string) (interface{}, error) {
	r := c.pool.Get()
	defer r.Close()

	return r.Do("set", key, value)
}

// GetString returns the string value stored with the given key.
//
// If the key doesn't exist, an error is returned.
func (c Cache) GetString(key string) (string, error) {
	r := c.pool.Get()
	defer r.Close()

	return redis.String(r.Do("get", key))
}

// PutMarshaled stores a json marshalled value with the given key.
func (c Cache) PutMarshaled(key string, value interface{}) (interface{}, error) {
	// Marshal to JSON
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Store in the cache
	return c.PutString(key, string(bytes[:]))
}

// GetMarshaled retrieves an item from the cache with the specified key,
// and un-marshals it from JSON to the value provided.
//
// If they key doesn't exist, an error is returned.
func (c Cache) GetMarshaled(key string, v interface{}) error {
	cached, err := c.GetString(key)
	if err != nil {
		return err
	}

	if len(cached) > 0 {
		if err := json.Unmarshal([]byte(cached), v); err != nil {
			return err
		}
	}
	return nil
}

// Delete removes an item from the cache by it's key.
func (c Cache) Delete(key string) error {
	r := c.pool.Get()
	defer r.Close()

	if _, err := r.Do("del", key); err != nil {
		return err
	}
	return nil
}

// Expire sets the time for a key to expire in seconds.
func (c Cache) Expire(key string, seconds time.Duration) error {
	r := c.pool.Get()
	defer r.Close()

	if _, err := r.Do("expire", key, seconds); err != nil {
		return err
	}
	return nil
}

// Lock attempts to put a lock on the key for a specified duration (in milliseconds).
// If the lock was successfully acquired, true will be returned.
//
// Note: The value provided can be anything, so long as it's unique. The value will then be used when
// attempting to Unlock, and will only work if the value matches. It's important that each instance that tries
// to perform a Lock have it's own unique key so that you don't unlock another instances lock!
func (c Cache) Lock(key, value string, timeoutMs int) (bool, error) {
	r := c.pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, lockScript)
	res, err := cmd.Do(r, key, value, timeoutMs)
	if err != nil {
		return false, err
	}

	return res == "OK", nil
}

// Unlock attempts to remove the lock on a key so long as the value matches.
// If the lock cannot be removed, either because the key has already expired or
// because the value was incorrect, an error will be returned.
func (c Cache) Unlock(key, value string) error {
	r := c.pool.Get()
	defer r.Close()

	cmd := redis.NewScript(1, unlockScript)
	if res, err := redis.Int(cmd.Do(r, key, value)); err != nil {
		return err
	} else if res != 1 {
		return ErrCantUnlock
	}

	return nil
}
