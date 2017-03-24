package convert

import (
	"testing"
)

func TestStringSliceToIntSlice(t *testing.T) {
	// Test empty
	empty := make([]string, 0, 0)
	if res, err := StringSliceToIntSlice(empty); err != nil {
		t.Error(err)
	} else if len(res) > 0 {
		t.Fatalf("Expected empty string slice to return empty int slice[]")
	}

	// Test proper use case
	proper := []string{"1", "-1", "0"}
	if res, err := StringSliceToIntSlice(proper); err != nil {
		t.Error(err)
	} else if len(res) != len(proper) {
		t.Fatalf("Incorrect slice size returned. got=%v, expected=%v", len(res), len(proper))
	} else if res[0] != 1 || res[1] != -1 || res[2] != 0 {
		t.Fatalf("Incorrect slice returned: %v", res)
	}

	// Test an invalid element
	invalid := []string{"1", "invalid"}
	if _, err := StringSliceToIntSlice(invalid); err == nil {
		t.Fatal("Expected error!")
	}
}

func TestIntSliceToStringSlice(t *testing.T) {
	// Test empty
	empty := make([]int, 0, 0)
	if len(IntSliceToStringSlice(empty)) != 0 {
		t.Fatal("Expected empty int slice to return empty string slice[]")
	}

	// Test proper use case
	proper := []int{1, 2, 3}
	strings := IntSliceToStringSlice(proper)
	if strings[0] != "1" || strings[1] != "2" || strings[2] != "3" {
		t.Fatalf("Unexpected result: %v", strings)
	}
}

func TestSliceToStringSlice(t *testing.T) {
	// Test empty
	empty := make([]interface{}, 0, 0)
	if len(SliceToStringSlice(empty)) != 0 {
		t.Fatalf("Expected empty slice to return empty string slice: %v", SliceToStringSlice(empty))
	}

	// Valid
	proper := []interface{}{1, "test"}
	strings := SliceToStringSlice(proper)
	if strings[0] != "1" || strings[1] != "test" {
		t.Fatalf("Unexpected Result: %v", strings)
	}
}
