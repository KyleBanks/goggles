package gonamo

import (
	"testing"
)

func TestAwsStringOrNil(t *testing.T) {
	tests := []struct {
		val       string
		expectNil bool
	}{
		{"value", false},
		{" ", false},
		{"", true},
	}

	for _, tt := range tests {
		s := AwsStringOrNil(tt.val)
		if tt.expectNil && s != nil {
			t.Fatalf("Expected string to be nil, got=%v", s)
		} else if !tt.expectNil && s == nil {
			t.Fatal("Expected string to not be nil")
		} else if !tt.expectNil && *s != tt.val {
			t.Fatalf("Unexpected value for string, expected=%v, got=%v", tt.val, *s)
		}
	}
}

func TestAttributeValue(t *testing.T) {
	{
		out := AttributeValue(StringType, nil)

		if out != nil {
			t.Fatalf("Expected nil output when input value is nil, got=%v", *out)
		}
	}

	{
		val := "string"
		out := AttributeValue(StringType, val)
		if *out.S != val {
			t.Fatalf("Unexpected AttributeValue, expected out.S=%v, got=%v", val, *out.S)
		}
	}

	{
		val := []string{"string1", "string2"}
		out := AttributeValue(StringArrayType, val)

		if len(val) != len(out.SS) {
			t.Fatalf("Output length mismatch, expected=%v, got=%v", len(val), len(out.SS))
		}

		for i := 0; i < len(val); i++ {
			if *out.SS[i] != val[i] {
				t.Fatalf("Unexpected AttributeValue, expected out.SS[%v]=%v, got=%v", i, val, *out.SS[i])
			}
		}
	}

	{
		val := true
		out := AttributeValue(BoolType, val)
		if *out.BOOL != val {
			t.Fatalf("Unexpected AttributeValue, expected out.BOOL=%v, got=%v", val, *out.BOOL)
		}
	}
}

func TestAttributeDefinition(t *testing.T) {
	tests := []struct {
		name          string
		attributeType AttributeType
	}{
		{"string", StringType},
		{"number", NumberType},
		{"bool", BoolType},
	}

	for _, tt := range tests {
		def := attributeDefinition(tt.name, tt.attributeType)
		if def == nil {
			t.Fatal("Expected attributeDefition to not be nil")
		}

		if *def.AttributeName != tt.name {
			t.Fatalf("Unexpected AttributeName, expected=%v, got=%v", tt.name, *def.AttributeName)
		}
		if *def.AttributeType != string(tt.attributeType) {
			t.Fatalf("Unexpected AttributeType, expected=%v, got=%v", tt.attributeType, *def.AttributeType)
		}
	}
}

func TestKeySchemaElement(t *testing.T) {
	tests := []struct {
		name    string
		keyType KeyType
	}{
		{"hashKey", HashKey},
		{"rangeKey", RangeKey},
	}

	for _, tt := range tests {
		key := keySchemaElement(tt.name, tt.keyType)
		if key == nil {
			t.Fatal("Expected keySchemaElement to not be nil")
		}

		if *key.AttributeName != tt.name {
			t.Fatalf("Unexpected AttributeName, expected=%v, got=%v", tt.name, *key.AttributeName)
		}
		if *key.KeyType != string(tt.keyType) {
			t.Fatalf("Unexpected KeyType, expected=%v, got=%v", tt.keyType, *key.KeyType)
		}
	}
}
