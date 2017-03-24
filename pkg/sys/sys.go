package sys

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/skratchdot/open-golang/open"
)

const (
	srcDirName      = "src"
	gopathEnv       = "GOPATH"
	gopathSeperator = ":"
)

var (
	cmdOpenTerminal = []string{"open", "-a", "Terminal"}

	defaultGoPath = os.ExpandEnv("$HOME/go")
)

// OpenFileExplorer opens the system file explorer application to the
// specified package.
func OpenFileExplorer(pkg string) {
	open.Run(AbsPath(pkg))
}

// OpenTerminal opens the system terminal (command line) application to the
// specified package.
func OpenTerminal(pkg string) {
	args := append(cmdOpenTerminal[1:], AbsPath(pkg))
	DefaultRunner.Run(cmdOpenTerminal[0], args...)
}

// OpenBrowser opens the default browser to the url provided.
func OpenBrowser(url string) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	open.Run(url)
}

// AbsPath returns the absolute path to a package from it's name.
func AbsPath(pkg string) string {
	srcDirs := Srcdir()
	for _, d := range srcDirs {
		info, err := os.Stat(filepath.Join(d, pkg))
		if err != nil || !info.IsDir() {
			continue

		}

		return filepath.Join(d, pkg)
	}

	return ""
}

// Srcdir returns the source directory(s) for go packages.
func Srcdir() []string {
	gopaths := Gopath()
	srcDirs := make([]string, len(gopaths), len(gopaths))
	for i, p := range gopaths {
		srcDirs[i] = filepath.Join(p, srcDirName)
	}
	return srcDirs
}

// Gopath returns the system GOPATH(s), defaulting to $HOME/go
// if the environment variable is not set.
func Gopath() []string {
	return strings.Split(RawGopath(), gopathSeperator)
}

// RawGopath returns the system GOPATH, defaulting to $HOME/go
// if the environment variable is not set, as a string regardless
// of how many GOPATHs are actually set.
func RawGopath() string {
	gopath := os.Getenv(gopathEnv)
	if len(gopath) == 0 {
		gopath = defaultGoPath
	}

	return gopath
}

// SetGopath sets the $GOPATH environment variable to the value provided.
func SetGopath(gopath string) {
	os.Setenv(gopathEnv, gopath)
}
