// +build !windows,!darwin

package sys

import (
	"fmt"
	"os/exec"
)

func init() {
	_, err := exec.LookPath("xterm")
	CanOpenTerminal = err == nil
}

func openTerminal(pkg string) {
	DefaultRunner.Run("xterm", "-e", fmt.Sprintf(`cd "%v; bash"`, AbsPath(pkg)))
}
