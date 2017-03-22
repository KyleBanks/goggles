package goggles

import (
	"testing"

	"github.com/KyleBanks/goggles/pkg/sys"
)

func Test_cleanPath(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"/foo/bar", "foo/bar"},
		{sys.Srcdir() + "/foo/bar", "foo/bar"},
	}

	for idx, tt := range tests {
		if out := cleanPath(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected output, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}

func Test_ignorePkg(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"/foo/bar/vendor/baz", true},
		{"/foo/bar/.git/baz", true},
		{"/foo/bar/testdata/baz", true},
		{"/foo/bar/baz", false},
	}

	for idx, tt := range tests {
		if out := ignorePkg(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected output, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}

func Test_repo(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"github.com/foo/bar", "github.com/foo/bar"},
		{"github.com/foo/bar/baz/etc", "github.com/foo/bar"},
		{"github.com/foo", ""},
		{"github.com", ""},
	}

	for idx, tt := range tests {
		if out := repo(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected repo, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}
