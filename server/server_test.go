package server

import (
	"net/http"
	"net/url"
	"testing"
)

func Test_New(t *testing.T) {
	m := New(nil)

	// Try some sample routes known to exist, ensuring they are registered
	tests := []struct {
		url    string
		expect string
	}{
		{"/api/pkg/details?name=test", "/api/pkg/details"},
		{"/static/css/main.css", "/"}, // Static assets are all resolved by "/"
	}

	for idx, tt := range tests {
		url, _ := url.Parse(tt.url)
		r := http.Request{
			Method: "GET",
			URL:    url,
		}
		_, p := m.Handler(&r)
		if p != tt.expect {
			t.Fatalf("[%v] Unexpected pattern, expected=%v, got=%v", idx, tt.expect, p)
		}
	}
}
