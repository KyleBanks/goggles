package goggles

import (
	"strings"

	"github.com/KyleBanks/goggles/pkg/sys"
)

var (
	ignorePaths = []string{".git", "/vendor/", "/testdata/"}
)

// cleanPath sanitizes a package path by removing the $GOPATH/src portion
// and any prefixed slashes.
func cleanPath(path string) string {
	path = strings.Replace(path, sys.Srcdir(), "", 1)
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
