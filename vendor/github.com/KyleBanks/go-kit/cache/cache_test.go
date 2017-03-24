package cache

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var (
	ValidHost = "localhost:6379"

	TestKey   = "Test:" + strconv.Itoa(int(time.Now().Unix()))
	TestValue = "Value:" + strconv.Itoa(int(time.Now().Unix()))
)

func TestNew(t *testing.T) {
	cache := New(ValidHost)
	if cache == nil {
		t.Error("Nil cache returned!")
	}
}

func TestCache_PutString(t *testing.T) {
	cache := New(ValidHost)

	if ok, err := cache.PutString(TestKey, TestValue); err != nil {
		t.Error(err)
	} else if ok != "OK" {
		t.Error("Unexpected cache response:", ok)
	}
}

func TestCache_GetString(t *testing.T) {
	cache := New(ValidHost)

	if res, err := cache.GetString(TestKey); err != nil {
		t.Error(err)
	} else if res != TestValue {
		t.Error("Unexpected result:", res)
	}
}

func TestCache_Delete(t *testing.T) {
	cache := New(ValidHost)

	// Delete should not return an error
	if err := cache.Delete(TestKey); err != nil {
		t.Error(err)
	}

	// Calling again, on a deleted key, should still not fail
	if err := cache.Delete(TestKey); err != nil {
		t.Error(err)
	}
}

func TestCache_GetString_DeletedKey(t *testing.T) {
	cache := New(ValidHost)

	if _, err := cache.GetString(TestKey); err == nil {
		t.Error("Expected error getting deleted key!")
	}
}

func TestCache_Lock(t *testing.T) {
	cache := New(ValidHost)

	key := fmt.Sprintf("testLock:%v", time.Now().Unix())
	value := "avalue"

	// Base test
	if locked, err := cache.Lock(key, value, 1000); err != nil {
		t.Fatal(err)
	} else if !locked {
		t.Fatal("Expected valid lock to return true")
	}

	// Try to lock the same key
	if locked, err := cache.Lock(key, value, 1000); err != nil {
		t.Fatal(err)
	} else if locked {
		t.Fatal("Expected invalid lock to return false")
	}

	// Wait
	time.Sleep(time.Millisecond * 1100)

	// Try again
	if locked, err := cache.Lock(key, value, 1000); err != nil {
		t.Fatal(err)
	} else if !locked {
		t.Fatal("Expected valid lock to return true")
	}
}

func TestCache_Unlock(t *testing.T) {
	cache := New(ValidHost)

	key := fmt.Sprintf("testUnlock:%v", time.Now().Unix())
	value := "avalue"

	cache.Lock(key, value, 1000)

	// Bad key
	if err := cache.Unlock("badkey", value); err != ErrCantUnlock {
		t.Fatal(err)
	}

	// Bad value
	if err := cache.Unlock(key, "badvalue"); err != ErrCantUnlock {
		t.Fatal(err)
	}

	// Valid
	if err := cache.Unlock(key, value); err != nil {
		t.Fatal(err)
	}
}
