package sys

import (
	"testing"
)

func TestCmdRunner_Run(t *testing.T) {
	var r CmdRunner

	// Expect an error
	_, err := r.Run("fake command", "this", "will", "fail")
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	// Expect success
	out, err := r.Run("echo", "hi")
	if err != nil {
		t.Fatal(err)
	} else if string(out) != "hi\n" {
		t.Fatalf("Unexpected output, expected=%v, got=%v", "hi", string(out))
	}
}
