package sys

import (
	"os"
	"path/filepath"
)

const (
	cmdOpenFileExplorer = "open"
	srcDirName          = "src"
	gopathEnv           = "GOPATH"
)

// OpenFileExplorer opens the system file explorer application to the
// specified package.
func OpenFileExplorer(name string) {
	Runner.Run(cmdOpenFileExplorer, AbsPath(name))
}

// AbsPath returns the absolute path to a package from it's name.
func AbsPath(name string) string {
	return filepath.Join(Srcdir(), name)
}

// Srcdir returns the source directory for go packages.
func Srcdir() string {
	return filepath.Join(Gopath(), srcDirName)
}

// Gopath returns the $GOPATH environment variable.
func Gopath() string {
	return os.Getenv(gopathEnv)
}
