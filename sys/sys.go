package sys

import (
	"os/exec"

	"github.com/KyleBanks/goggles/pkg"
)

// Default is the default System.
var Default System

// System is a type that can interact with the host OS.
type System struct{}

// OpenFileExplorer opens the system file explorer application to the
// specified package.
func (System) OpenFileExplorer(name string) {
	c := exec.Command("open", pkg.AbsPath(name))
	c.Run()
}
