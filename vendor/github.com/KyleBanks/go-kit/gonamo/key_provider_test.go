package gonamo

import (
	"testing"
)

func TestHashKeyDefinitionHashName(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"hashName"},
	}

	for _, tt := range tests {
		h := HashKeyDefinition{HashName: tt.input}

		if h.hashName() != tt.input {
			t.Fatalf("Unexpected hashName, expected=%v, got=%v", tt.input, h.hashName())
		}
	}
}

func TestHashKeyDefinitionHashType(t *testing.T) {
	tests := []struct {
		input AttributeType
	}{
		{StringType},
		{NumberType},
	}

	for _, tt := range tests {
		h := HashKeyDefinition{HashType: tt.input}

		if h.hashType() != tt.input {
			t.Fatalf("Unexpected hashName, expected=%v, got=%v", tt.input, h.hashType())
		}
	}
}

func TestHashKeyDefinitionHasRange(t *testing.T) {
	h := HashKeyDefinition{}

	if h.hasRange() {
		t.Fatalf("Unexpected hasRange, expected=%v, got=%v", false, h.hasRange())
	}
}

func TestHashKeyDefinitionRangeName(t *testing.T) {
	h := HashKeyDefinition{}

	if h.rangeName() != "" {
		t.Fatalf("Unexpected rangeName, expected=%v, got=%v", "", h.rangeName())
	}
}

func TestHashKeyDefinitionRangeType(t *testing.T) {
	h := HashKeyDefinition{}

	if h.rangeType() != "" {
		t.Fatalf("Unexpected rangeType, expected=%v, got=%v", "", h.rangeType())
	}
}

func TestHashRangeKeyDefinitionHashName(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"hashName"},
	}

	for _, tt := range tests {
		h := HashRangeKeyDefinition{HashName: tt.input}

		if h.hashName() != tt.input {
			t.Fatalf("Unexpected hashName, expected=%v, got=%v", tt.input, h.hashName())
		}
	}
}

func TestHashRangeKeyDefinitionHashType(t *testing.T) {
	tests := []struct {
		input AttributeType
	}{
		{StringType},
		{NumberType},
	}

	for _, tt := range tests {
		h := HashRangeKeyDefinition{HashType: tt.input}

		if h.hashType() != tt.input {
			t.Fatalf("Unexpected hashName, expected=%v, got=%v", tt.input, h.hashType())
		}
	}
}

func TestHashRangeKeyDefinitionHasRange(t *testing.T) {
	h := HashRangeKeyDefinition{}

	if !h.hasRange() {
		t.Fatalf("Unexpected hasRange, expected=%v, got=%v", true, h.hasRange())
	}
}

func TestHashRangeKeyDefinitionRangeName(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"rangeName"},
	}

	for _, tt := range tests {
		h := HashRangeKeyDefinition{RangeName: tt.input}

		if h.rangeName() != tt.input {
			t.Fatalf("Unexpected rangeName, expected=%v, got=%v", tt.input, h.rangeName())
		}
	}
}

func TestHashRangeKeyDefinitionRangeType(t *testing.T) {
	tests := []struct {
		input AttributeType
	}{
		{StringType},
		{NumberType},
	}

	for _, tt := range tests {
		h := HashRangeKeyDefinition{RangeType: tt.input}

		if h.rangeType() != tt.input {
			t.Fatalf("Unexpected rangeType, expected=%v, got=%v", tt.input, h.rangeType())
		}
	}
}
