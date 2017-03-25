package resolver

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
		{sys.Srcdir()[0] + "/foo/bar", "foo/bar"},
	}

	for idx, tt := range tests {
		if out := cleanPath(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected output, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}

func Test_ignore(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"github.com/foo/bar/vendor/baz", true},
		{"github.com/foo/bar/.git/baz", true},
		{"github.com/foo/bar/testdata/baz", true},
		{"github.com/foo/bar/baz", false},
	}

	for idx, tt := range tests {
		if out := ignore(tt.in); out != tt.out {
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

func Test_docs(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		// Code blocks
		{"   example", "    example"},                                                 // 3 -> 4 spaces
		{"        example", "        example"},                                        // >= 4 stays the same
		{"  example", "  example"},                                                    // < 3 stays the same
		{"   example\n\texample", "    example\n\texample"},                           // multiple lines
		{"\texample", "\texample"},                                                    // tabs
		{"\texample\n   example\n    example", "\texample\n    example\n    example"}, // tabs, three spaces and four spaces
	}

	for idx, tt := range tests {
		if out := docs(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected doc, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}
