package sys

import (
	"os/exec"
)

// Runner is the default command runner used by the sys package.
var Runner ICmdRunner = CmdRunner{}

// ICmdRunner defines a type that can run system commands.
type ICmdRunner interface {
	Run(string, ...string) error
}

// CmdRunner runs system commands.
type CmdRunner struct{}

// Run executes a system command.
func (CmdRunner) Run(cmd string, args ...string) error {
	return exec.Command(cmd, args...).Run()
}
