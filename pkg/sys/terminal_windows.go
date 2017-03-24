package sys

import (
	"fmt"
)

func init() {
	CanOpenTerminal = true
}

func openTerminal(pkg string) {
	DefaultRunner.Run("cmd", "/K", fmt.Sprintf(`"cd /d %v"`, AbsPath(pkg)))
}
