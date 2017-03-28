package release

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetLatest(t *testing.T) {
	owner := "foo"
	repo := "bar"
	expect := "v1.0.0"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		if query.Get("owner") != owner || query.Get("repo") != repo {
			t.Fatalf("Unexpected owner/repo, expected=%v/%v, got=%v/%v", owner, repo, query.Get("owner"), query.Get("repo"))
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[{"name": "`+expect+`"}]`)
	}))
	defer ts.Close()
	endpoint = ts.URL + "?owner=%v&repo=%v"

	out, err := GetLatest(owner, repo)
	if err != nil {
		t.Fatal(err)
	} else if out != expect {
		t.Fatalf("Unexpected version, expected=%v, got=%v", expect, out)
	}
}

func Test_IsLatest(t *testing.T) {
	version := "v1.0.0"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `[{"name": "`+version+`"}]`)
	}))
	defer ts.Close()
	endpoint = ts.URL + "?repo=%v&name=%v"

	// positive case
	{
		isLatest, v, err := IsLatest("foo", "bar", version)
		if err != nil {
			t.Fatal(err)
		}

		if !isLatest {
			t.Fatalf("Expected isLatest to be true, got version=%v", v)
		}
	}

	// negative case
	{
		isLatest, v, err := IsLatest("KyleBanks", "goggles", version+".789")
		if err != nil {
			t.Fatal(err)
		}

		if isLatest {
			t.Fatalf("Expected isLatest to be false, got version=%v", v)
		}
	}
}
