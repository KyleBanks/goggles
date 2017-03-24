package sys

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	srcDirName = "src"
	gopathEnv  = "GOPATH"
)

var (
	cmdOpenFileExplorer = []string{"open"}
	cmdOpenTerminal     = []string{"open", "-a", "Terminal"}
	cmdOpenBrowser      = []string{"open"}

	defaultGoPath = os.ExpandEnv("$HOME/go")
)

// OpenFileExplorer opens the system file explorer application to the
// specified package.
func OpenFileExplorer(pkg string) {
	DefaultRunner.Run(cmdOpenFileExplorer[0], AbsPath(pkg))
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

	DefaultRunner.Run(cmdOpenBrowser[0], url)
}

// AbsPath returns the absolute path to a package from it's name.
func AbsPath(pkg string) string {
	return filepath.Join(Srcdir(), pkg)
}

// Srcdir returns the source directory for go packages.
func Srcdir() string {
	return filepath.Join(Gopath(), srcDirName)
}

// Gopath returns the $GOPATH environment variable, defaulting to $HOME/go
// if the environment variable is not set.
func Gopath() string {
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
