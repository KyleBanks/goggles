package goggles

import (
	"os/exec"
	"path/filepath"
)

// OpenFileExplorer opens the system file explorer application to the
// specified package.
func OpenFileExplorer(name string) {
	c := exec.Command("open", filepath.Join(srcdir(), name))
	c.Run()
}
