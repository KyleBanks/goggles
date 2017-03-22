package sys

import (
	"os/exec"
)

// DefaultRunner is the default command runner used by the sys package.
var DefaultRunner Runner = CmdRunner{}

// Runner defines a type that can run system commands.
type Runner interface {
	Run(string, ...string) ([]byte, error)
}

// CmdRunner runs system commands.
type CmdRunner struct{}

// Run executes a system command.
func (CmdRunner) Run(cmd string, args ...string) ([]byte, error) {
	return exec.Command(cmd, args...).CombinedOutput()
}
