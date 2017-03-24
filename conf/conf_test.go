package conf

import (
	"testing"
)

type mockSaveLoader struct {
	saveFn func(interface{}) error
	loadFn func(interface{}) error
}

func (m mockSaveLoader) Save(v interface{}) error { return m.saveFn(v) }
func (m mockSaveLoader) Load(v interface{}) error { return m.loadFn(v) }

func Test_Get(t *testing.T) {
	expect := "/foo/bar"
	var m mockSaveLoader
	m.loadFn = func(v interface{}) error {
		if _, ok := v.(*Config); !ok {
			t.Fatalf("Unexpected Type loaded, expected=%v, got=%v", &Config{}, v)
		}

		v.(*Config).Gopath = expect
		return nil
	}
	store = m

	c := Get()
	if c.Gopath != expect {
		t.Fatalf("Expected Config returned, expected=%v, got=%v", expect, c)
	}
}

func Test_Save(t *testing.T) {
	expect := "/foo/bar"

	var called bool
	var m mockSaveLoader
	m.saveFn = func(v interface{}) error {
		if v.(*Config).Gopath != expect {
			t.Fatalf("Expected Config returned, expected=%v, got=%v", expect, v.(*Config).Gopath)
		}
		called = true
		return nil
	}
	store = m

	c := Config{Gopath: expect}
	if err := Save(&c); err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Expected Save to be called")
	}
}
