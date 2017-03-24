package cmd

import (
	"reflect"
	"testing"

	"github.com/KyleBanks/goggles/conf"
	"github.com/KyleBanks/goggles/pkg/sys"
)

func Test_initConfig(t *testing.T) {
	tests := []struct {
		in     string
		expect []string
	}{
		{"", sys.Gopath()},
		{"/foo/bar", []string{"/foo/bar"}},
		{"/foo/bar:/foo/bar/baz", []string{"/foo/bar", "/foo/bar/baz"}},
	}

	for idx, tt := range tests {
		defaultProvider = &mockProvider{
			preferencesFn: func() *conf.Config {
				return &conf.Config{Gopath: tt.in}
			},
		}

		initConfig()

		if out := sys.Gopath(); !reflect.DeepEqual(out, tt.expect) {
			t.Fatalf("[%v] Unexpected Gopath, expected=%v, got=%v", idx, tt.expect, out)
		}
	}
}

func Test_OpenLinks(t *testing.T) {
	tests := []struct {
		expect string
		fn     func()
	}{
		{aboutURL, OpenAbout},
		{thanksURL, OpenThanks},
	}

	for _, tt := range tests {
		var called bool
		defaultProvider = &mockProvider{
			OpenBrowserFn: func(s string) {
				if s != tt.expect {
					t.Fatalf("Unexpected URL, expected=%v, got=%v", tt.expect, s)
				}
				called = true
			},
		}

		tt.fn()
		if !called {
			t.Fatal("Expected OpenBrowser to be called")
		}
	}
}
