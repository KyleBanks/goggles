package sys

import (
	"testing"
)

func TestCmdRunner_Run(t *testing.T) {
	var r CmdRunner

	// Expect an error
	err := r.Run("fake command", "this", "will", "fail")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Expect success
	err = r.Run("echo", "hi")
	if err != nil {
		t.Fatal(err)
	}
}
