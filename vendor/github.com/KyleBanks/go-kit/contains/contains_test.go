package contains

import (
	"testing"
)

func TestInt(t *testing.T) {
	// Negative cases
	if Int(0, []int{}) {
		t.Fatal("Expected value not to be in empty slice")
	} else if Int(0, []int{1, 2, 3}) {
		t.Fatal("Expected value not to be in slice")
	}

	// Positive Cases
	if !Int(0, []int{1, 2, 3, 0}) {
		t.Fatal("Expected value to be in slice")
	}
}

func TestUint(t *testing.T) {
	// Negative cases
	if Uint(uint(0), []uint{}) {
		t.Fatal("Expected value not to be in empty slice")
	} else if Uint(uint(0), []uint{uint(1), uint(2), uint(3)}) {
		t.Fatal("Expected value not to be in slice")
	}

	// Positive Cases
	if !Uint(0, []uint{uint(1), uint(2), uint(3), uint(0)}) {
		t.Fatal("Expected value to be in slice")
	}
}
