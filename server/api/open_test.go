package api

import (
	"net/http/httptest"
	"testing"
)

func Test_openFileExplorer(t *testing.T) {
	m := setup()
	expect := "github.com/foo/bar"
	var called bool
	m.OpenFileExplorerFn = func(name string) {
		called = true

		if name != expect {
			t.Fatalf("Unexpected name, expected=%v, got=%v", expect, name)
		}
	}

	r := httptest.NewRequest("GET", "/?name="+expect, nil)
	w := httptest.NewRecorder()

	openFileExplorer(w, r)
	if !called {
		t.Fatal("Expected OpenFileExplorer to be called")
	}

	var out map[string]string
	validateResponse(t, w, &out)
}
