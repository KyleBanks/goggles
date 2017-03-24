package sys

func init() {
	CanOpenTerminal = true
}

func openTerminal(pkg string) {
	DefaultRunner.Run("open", "-a", "Terminal", AbsPath(pkg))
}
