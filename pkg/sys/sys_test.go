package sys

import (
	"os"
	"path/filepath"
	"testing"
)

type mockRunner struct {
	runFn func(string, ...string) error
}

func (m mockRunner) Run(cmd string, args ...string) error { return m.runFn(cmd, args...) }

func Test_OpenFileExplorer(t *testing.T) {
	expect := []string{"/foo/bar/gopath", "src", "github.com/foo/bar"}
	os.Setenv("GOPATH", expect[0])

	var gotCmd string
	var gotPath string
	Runner = mockRunner{
		runFn: func(cmd string, args ...string) error {
			gotCmd = cmd
			gotPath = args[0]
			return nil
		},
	}

	OpenFileExplorer(expect[2])

	if gotCmd != cmdOpenFileExplorer {
		t.Fatalf("Unexpected cmd, expected=%v, got=%v", cmdOpenFileExplorer, gotCmd)
	} else if gotPath != filepath.Join(expect...) {
		t.Fatalf("Unexpected path, expected=%v, got=%v", filepath.Join(expect...), gotPath)
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
	expect := "/foo/bar/path"
	os.Setenv("GOPATH", expect)

	if Gopath() != expect {
		t.Fatalf("Unexpected Gopath, expected=%v, got=%v", expect, Gopath())
	}
}
