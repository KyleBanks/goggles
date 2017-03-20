package pkg

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	ignorePaths = []string{".git", "/vendor/", "/testdata/"}
)

// AbsPath returns the absolute path to a package from it's name.
func AbsPath(name string) string {
	return filepath.Join(srcdir(), name)
}

// cleanPath sanitizes a package path.
func cleanPath(path string) string {
	path = strings.Replace(path, srcdir(), "", 1)
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return path
}

// ignorePkg checks if the path provided should be ignored.
func ignorePkg(path string) bool {
	if len(path) == 0 {
		return true
	}

	for _, s := range ignorePaths {
		if strings.Contains(path, s) {
			return true
		}
	}

	return false
}

// srcdir returns the source directory for go packages.
func srcdir() string {
	return filepath.Join(gopath(), "src")
}

// gopath returns the $GOPATH environment variable.
func gopath() string {
	return os.Getenv("GOPATH")
}
