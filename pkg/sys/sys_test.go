package sys

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type mockRunner struct {
	runFn func(string, ...string) ([]byte, error)
}

func (m mockRunner) Run(cmd string, args ...string) ([]byte, error) { return m.runFn(cmd, args...) }

func Test_OpenTerminal(t *testing.T) {
	expect := []string{os.ExpandEnv("$GOPATH"), "src", "github.com/KyleBanks/goggles"}

	var gotCmd string
	var gotArgs []string
	DefaultRunner = mockRunner{
		runFn: func(cmd string, args ...string) ([]byte, error) {
			gotCmd = cmd
			gotArgs = args
			return nil, nil
		},
	}

	OpenTerminal(expect[2])

	if gotCmd != cmdOpenTerminal[0] {
		t.Fatalf("Unexpected cmd, expected=%v, got=%v", cmdOpenTerminal[0], gotCmd)
	} else if gotArgs[0] != cmdOpenTerminal[1] || gotArgs[1] != cmdOpenTerminal[2] || gotArgs[2] != filepath.Join(expect...) {
		t.Fatalf("Unexpected args, expected=%v, %v, got=%v", cmdOpenTerminal[1:], filepath.Join(expect...), gotArgs)
	}
}

func Test_AbsPath(t *testing.T) {
	expect := []string{os.ExpandEnv("$GOPATH"), "src", "github.com/foo/bar"}

	if AbsPath(expect[2]) != filepath.Join(expect...) {
		t.Fatalf("Unexpected AbsPath, expected=%v, got=%v", filepath.Join(expect...), AbsPath(expect[2]))
	}
}

func Test_Srcdir(t *testing.T) {
	tests := []struct {
		env    string
		expect []string
	}{
		{"/foo/bar/path", []string{"/foo/bar/path/src"}},
		{"/foo/bar/path:/other/path", []string{"/foo/bar/path/src", "/other/path/src"}},
		{"", []string{defaultGoPath + "/src"}},
	}

	for idx, tt := range tests {
		os.Setenv("GOPATH", tt.env)

		if got := Srcdir(); !reflect.DeepEqual(got, tt.expect) {
			t.Fatalf("[%v] Unexpected Srcdir, expected=%v, got=%v", idx, tt.expect, got)
		}
	}
}

func Test_Gopath(t *testing.T) {
	tests := []struct {
		env    string
		expect []string
	}{
		{"/foo/bar/path", []string{"/foo/bar/path"}},
		{"/foo/bar/path:/other/path", []string{"/foo/bar/path", "/other/path"}},
		{"", []string{defaultGoPath}},
	}

	for idx, tt := range tests {
		os.Setenv("GOPATH", tt.env)

		if got := Gopath(); !reflect.DeepEqual(got, tt.expect) {
			t.Fatalf("[%v] Unexpected Gopath, expected=%v, got=%v", idx, tt.expect, got)
		}
	}
}

func Test_RawGopath(t *testing.T) {
	tests := []struct {
		env    string
		expect string
	}{
		{"/foo/bar/path", "/foo/bar/path"},
		{"/foo/bar/path:/other/path", "/foo/bar/path:/other/path"},
		{"", defaultGoPath},
	}

	for idx, tt := range tests {
		os.Setenv("GOPATH", tt.env)

		if got := RawGopath(); got != tt.expect {
			t.Fatalf("[%v] Unexpected RawGopath, expected=%v, got=%v", idx, tt.expect, got)
		}
	}
}

func Test_SetGopath(t *testing.T) {
	expect := "/foo/bar/custom/path"
	SetGopath(expect)

	if os.Getenv("GOPATH") != expect {
		t.Fatalf("Unexpected Gopath, expected=%v, got=%v", expect, os.Getenv("GOPATH"))
	}
}
