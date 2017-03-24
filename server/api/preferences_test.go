package api

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/KyleBanks/goggles/conf"
)

func Test_getPreferences(t *testing.T) {
	expect := &conf.Config{
		Gopath: "/foo/bar/path",
	}

	m := setup()
	m.preferencesFn = func() *conf.Config {
		return expect
	}

	r := httptest.NewRequest("GET", "/?gopath="+expect.Gopath, nil)
	w := httptest.NewRecorder()
	getPreferences(w, r)

	var out conf.Config
	validateResponse(t, w, &out)

	if !reflect.DeepEqual(&out, expect) {
		t.Fatalf("Unexpected Config, expected=%v, got=%v", expect, out)
	}
}

func Test_updatePreferences(t *testing.T) {
	expect := &conf.Config{
		Gopath: "/foo/bar/path",
	}

	var called bool
	m := setup()
	m.updatePreferencesFn = func(c *conf.Config) {
		if !reflect.DeepEqual(c, expect) {
			t.Fatalf("Unexpected Config, expected=%v, got=%v", expect, c)
		}

		called = true
	}

	r := httptest.NewRequest("GET", "/?gopath="+expect.Gopath, nil)
	w := httptest.NewRecorder()
	updatePreferences(w, r)

	var out map[string]string
	validateResponse(t, w, &out)

	if !called {
		t.Fatal("Expected UpdatePreferences to be called")
	}
}
