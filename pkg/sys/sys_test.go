package sys

import (
	"os"
	"path/filepath"
	"testing"
)

type mockRunner struct {
	runFn func(string, ...string) ([]byte, error)
}

func (m mockRunner) Run(cmd string, args ...string) ([]byte, error) { return m.runFn(cmd, args...) }

func Test_OpenFileExplorer(t *testing.T) {
	expect := []string{"/foo/bar/gopath", "src", "github.com/foo/bar"}
	os.Setenv("GOPATH", expect[0])

	var gotCmd string
	var gotPath string
	DefaultRunner = mockRunner{
		runFn: func(cmd string, args ...string) ([]byte, error) {
			gotCmd = cmd
			gotPath = args[0]
			return nil, nil
		},
	}

	OpenFileExplorer(expect[2])

	if gotCmd != cmdOpenFileExplorer[0] {
		t.Fatalf("Unexpected cmd, expected=%v, got=%v", cmdOpenFileExplorer[0], gotCmd)
	} else if gotPath != filepath.Join(expect...) {
		t.Fatalf("Unexpected path, expected=%v, got=%v", filepath.Join(expect...), gotPath)
	}
}

func Test_OpenTerminal(t *testing.T) {
	expect := []string{"/foo/bar/gopath", "src", "github.com/foo/bar"}
	os.Setenv("GOPATH", expect[0])

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

func Test_OpenBrowser(t *testing.T) {
	tests := []struct {
		url    string
		expect string
	}{
		{"github.com/foo/bar", "http://github.com/foo/bar"},
		{"http://github.com/foo/bar", "http://github.com/foo/bar"},
		{"https://github.com/foo/bar", "https://github.com/foo/bar"},
	}

	for idx, tt := range tests {
		var gotCmd string
		var gotURL string
		DefaultRunner = mockRunner{
			runFn: func(cmd string, args ...string) ([]byte, error) {
				gotCmd = cmd
				gotURL = args[0]
				return nil, nil
			},
		}

		OpenBrowser(tt.url)

		if gotCmd != cmdOpenBrowser[0] {
			t.Fatalf("[%v] Unexpected cmd, expected=%v, got=%v", idx, cmdOpenBrowser[0], gotCmd)
		} else if gotURL != tt.expect {
			t.Fatalf("[%v] Unexpected url, expected=%v, got=%v", idx, tt.expect, gotURL)
		}
	}
}

func Test_AbsPath(t *testing.T) {
	expect := []string{"/foo/bar/gopath", "src", "github.com/foo/bar"}
	os.Setenv("GOPATH", expect[0])

	if AbsPath(expect[2]) != filepath.Join(expect...) {
		t.Fatalf("Unexpected AbsPath, expected=%v, got=%v", filepath.Join(expect...), AbsPath(expect[2]))
	}
}

func Test_Srcdir(t *testing.T) {
	expect := []string{"/foo/bar/gopath", "src"}
	os.Setenv("GOPATH", expect[0])

	if Srcdir() != filepath.Join(expect...) {
		t.Fatalf("Unexpected Srcdir, expected=%v, got=%v", filepath.Join(expect...), Srcdir())
	}
}

func Test_Gopath(t *testing.T) {
	// GOPATH avaiable
	expect := "/foo/bar/path"
	os.Setenv("GOPATH", expect)
	if Gopath() != expect {
		t.Fatalf("Unexpected Gopath, expected=%v, got=%v", expect, Gopath())
	}

	// Default
	expect = defaultGoPath
	os.Setenv("GOPATH", "")
	if Gopath() != expect {
		t.Fatalf("Unexpected Gopath, expected=%v, got=%v", expect, Gopath())
	}
}

func Test_SetGopath(t *testing.T) {
	expect := "/foo/bar/custom/path"
	SetGopath(expect)

	if os.Getenv("GOPATH") != expect {
		t.Fatalf("Unexpected Gopath, expected=%v, got=%v", expect, os.Getenv("GOPATH"))
	}
}
