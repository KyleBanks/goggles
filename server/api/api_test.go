package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/resolver"
)

type mockProvider struct {
	ListFn    func() ([]*resolver.Package, error)
	DetailsFn func(string) (*resolver.Package, error)

	OpenFileExplorerFn func(string)
	OpenTerminalFn     func(string)
	OpenBrowserFn      func(string)

	preferencesFn       func() *conf.Config
	updatePreferencesFn func(*conf.Config)
}

func (m *mockProvider) List() ([]*resolver.Package, error)          { return m.ListFn() }
func (m *mockProvider) Details(n string) (*resolver.Package, error) { return m.DetailsFn(n) }
func (m *mockProvider) OpenFileExplorer(n string)                   { m.OpenFileExplorerFn(n) }
func (m *mockProvider) OpenTerminal(n string)                       { m.OpenTerminalFn(n) }
func (m *mockProvider) OpenBrowser(n string)                        { m.OpenBrowserFn(n) }
func (m *mockProvider) Preferences() *conf.Config                   { return m.preferencesFn() }
func (m *mockProvider) UpdatePreferences(c *conf.Config)            { m.updatePreferencesFn(c) }

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
	mux := http.NewServeMux()
	Bind(expect, mux)

	if provider != expect {
		t.Fatalf("Unexpected provider, expected=%v, got=%v", expect, provider)
	}

	// Ensure routes are registered
	tests := []struct {
		url string
	}{
		{"/api/pkg/list"},
		{"/api/pkg/details"},
		{"/api/open/file-explorer"},
		{"/api/open/terminal"},
		{"/api/open/url"},
		{"/api/preferences/"},
		{"/api/preferences/update"},
	}

	for idx, tt := range tests {
		url, _ := url.Parse(tt.url)
		r := http.Request{
			Method: "GET",
			URL:    url,
		}

		_, p := mux.Handler(&r)
		if p != tt.url {
			t.Fatalf("[%v] Unexpected pattern, expected=%v, got=%v", idx, tt.url, p)
		}
	}
}

func Test_outputEmpty(t *testing.T) {
	var b bytes.Buffer
	outputEmpty(&b)

	if b.String() != "{}" {
		t.Fatalf("Unexpected output, expected=%v, got=%v", "{}", b.String())
	}
}
