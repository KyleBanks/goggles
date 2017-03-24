package goggles

import (
	"strings"

	"github.com/KyleBanks/goggles/pkg/sys"
)

var (
	ignorePaths = []string{".git", ".build", "/Godep/", "/vendor/", "/testdata/"}
)

// cleanPath sanitizes a package path by removing the $GOPATH/src portion
// and any prefixed slashes.
func cleanPath(path string) string {
	// Remove the first Srcdir prefix found
	for _, dir := range sys.Srcdir() {
		if !strings.HasPrefix(path, dir) {
			continue
		}

		path = strings.Replace(path, dir, "", 1)
		break
	}

	if strings.HasPrefix(path, "/") || strings.HasPrefix(path, "\\") {
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

// repo returns the repository of the package provided.
//
// For example, if "github.com/foo/bar/baz" is provided,
// "github.com/foo/bar" will be returned. If the repository
// cannot be determined, an empty string is returned.
func repo(pkg string) string {
	components := strings.Split(pkg, "/")
	if len(components) <= 2 {
		return ""
	}

	return strings.Join(components[0:3], "/")
}
