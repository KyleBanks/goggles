package job

import (
	"testing"
	"time"
)

// TODO: These tests are almost certainly prone to race conditions.
func TestRegister(t *testing.T) {
	tests := []struct {
		runImmediately bool
	}{
		{true},
		{false},
	}

	for _, tt := range tests {
		callCount := 0
		delay := time.Millisecond * 500
		j := Register(func() {
			callCount++
		}, delay, tt.runImmediately)

		if j == nil {
			t.Fatal("Unexpected nil response from Register")
		}

		// Allow a moment for the first run to (potentially) execute.
		time.Sleep(time.Millisecond * 100)

		if callCount > 0 && !tt.runImmediately {
			t.Fatalf("Unexpected callCount for runImmediately=false, expected=%v, got=%v", 0, callCount)
		} else if callCount == 0 && tt.runImmediately {
			t.Fatalf("Unexpected callCount for runImmediately=true, expected=%v, got=%v", 1, callCount)
		}

		// Reset callCount if runImmediately
		if tt.runImmediately {
			callCount--
		}

		for i := 1; i <= 2; i++ {
			time.Sleep(delay + time.Millisecond*100)

			if callCount != i {
				t.Fatalf("Unexpected callCount, expected=%v, got=%v", i, callCount)
			}
		}
	}
}

func TestStop(t *testing.T) {
	{
		// Should still runImmediately even if stop is immediately called
		callCount := 0
		j := Register(func() {
			callCount++
		}, time.Hour, true)

		j.Stop()

		// Allow a moment for the first run execute.
		time.Sleep(time.Millisecond * 100)

		if callCount == 0 {
			t.Fatal("Expected runImmediately to still execute first time even when Stop() is called.")
		}
	}

	{
		// Should stop executions after Stop is called
		callCount := 0
		delay := time.Millisecond * 500
		j := Register(func() {
			callCount++
		}, delay, false)

		// Wait for first execution
		time.Sleep(delay + time.Millisecond*100)
		if callCount != 1 {
			t.Fatalf("Unexpected callCount (before Stop), expected=%v, got=%v", 1, callCount)
		}

		j.Stop()

		// Wait for the second execution, which should never happen
		time.Sleep(delay + time.Millisecond*100)

		if callCount != 1 {
			t.Fatalf("Unexpected callCount (after Stop), expected=%v, got=%v", 1, callCount)
		}
	}
}
