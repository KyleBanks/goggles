package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KyleBanks/goggles/goggles"
)

type mockProvider struct {
	ListFn    func() ([]*goggles.Package, error)
	DetailsFn func(string) (*goggles.Package, error)

	OpenDevToolsFn     func()
	OpenFileExplorerFn func(string)
}

func (m *mockProvider) List() ([]*goggles.Package, error)          { return m.ListFn() }
func (m *mockProvider) Details(n string) (*goggles.Package, error) { return m.DetailsFn(n) }
func (m *mockProvider) OpenDevTools()                              { m.OpenDevToolsFn() }
func (m *mockProvider) OpenFileExplorer(n string)                  { m.OpenFileExplorerFn(n) }

func setup() *mockProvider {
	m := &mockProvider{}
	provider = m

	return m
}

func validateResponse(t *testing.T, w *httptest.ResponseRecorder, v interface{}) {
	if w.Code != 200 {
		t.Fatalf("Unexpected response code, expected=%v, got=%v", 200, w.Code)
	}

	err := json.NewDecoder(w.Body).Decode(v)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Bind(t *testing.T) {
	expect := &mockProvider{}
	Bind(expect, &http.ServeMux{})

	if provider != expect {
		t.Fatalf("Unexpected provider, expected=%v, got=%v", expect, provider)
	}
}

func Test_outputEmpty(t *testing.T) {
	var b bytes.Buffer
	outputEmpty(&b)

	if b.String() != "{}" {
		t.Fatalf("Unexpected output, expected=%v, got=%v", "{}", b.String())
	}
}
