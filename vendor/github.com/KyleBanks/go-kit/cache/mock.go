package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/KyleBanks/go-kit/log"
)

// Mock provides a mocked Cache implementation for testing.
type Mock struct {
	mu    *sync.Mutex
	cache map[string]string
}

// NewMock instantiates and returns a new Mock cache.
func NewMock() *Mock {
	return &Mock{
		mu:    &sync.Mutex{},
		cache: make(map[string]string),
	}
}

// PutString stores a simple key-value pair in the mock.
func (m Mock) PutString(key string, value string) (interface{}, error) {
	log.Info("Mock PutString:", key, value)
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[key] = value
	return nil, nil
}

// GetString returns the string value stored with the given key.
//
// If the key doesn't exist, an error is returned.
func (m Mock) GetString(key string) (string, error) {
	log.Info("Mock GetString:", key)
	m.mu.Lock()
	defer m.mu.Unlock()

	val, ok := m.cache[key]
	if !ok {
		return "", errors.New("Key not found")
	}

	return val, nil
}

// PutMarshaled stores a json marshalled value with the given key.
func (m Mock) PutMarshaled(key string, value interface{}) (interface{}, error) {
	// Marshal to JSON
	bytes, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}

	// Store in the cache
	return m.PutString(key, string(bytes[:]))
}

// GetMarshaled retrieves an item from the cache with the specified key,
// and un-marshals it from JSON to the value provided.
//
// If they key doesn't exist, an error is returned.
func (m Mock) GetMarshaled(key string, v interface{}) error {
	cached, err := m.GetString(key)
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

// Delete removes an item from the mock by it's key.
func (m Mock) Delete(key string) error {
	log.Info("Mock Delete:", key)
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.cache, key)
	return nil
}

// Expire sets the time for a key to expire in seconds.
func (m Mock) Expire(key string, seconds time.Duration) error {
	log.Info("Mock Expire:", key, seconds)

	go func() {
		time.Sleep(time.Second * time.Duration(seconds))
		m.Delete(key)
	}()
	return nil
}

// Lock attempts to put a lock on the key for a specified duration (in milliseconds).
func (m Mock) Lock(key, value string, durationMs int) (bool, error) {
	// TODO: Probably a better way to do this
	if _, err := m.GetString(key); err == nil {
		return false, errors.New("key already exists")
	}

	m.PutString(key, value)
	d, _ := time.ParseDuration(fmt.Sprintf("%vms", durationMs))
	m.Expire(key, d)
	return true, nil
}

// Unlock attempts to remove the lock on a key so long as the value matches.
func (m Mock) Unlock(key, value string) error {
	// TODO: Probably a better way to do this

	if val, err := m.GetString(key); err != nil {
		return err
	} else if val != value {
		return errors.New("value mismatch")
	}

	m.Delete(key)
	return nil
}
