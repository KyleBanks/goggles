package api

import (
	"net/http/httptest"
	"testing"
)

func Test_debug(t *testing.T) {
	m := setup()
	var called bool
	m.OpenDevToolsFn = func() {
		called = true
	}

	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	debug(w, r)
	if !called {
		t.Fatal("Expected OpenDevTools to be called")
	}

	var out map[string]string
	validateResponse(t, w, &out)
}
