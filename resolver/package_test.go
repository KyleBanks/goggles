package resolver

import (
	"testing"
)

func Test_NewPackage(t *testing.T) {
	path := "github.com/KyleBanks/goggles/resolver"
	p, err := NewPackage(path)
	if err != nil {
		t.Fatal(err)
	}

	if p.Name != path {
		t.Fatalf("Unexpected Name, expected=%v, got=%v", path, p.Name)
	}
}

func Test_cleanDoc(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		// Code blocks
		{"   example", "    example"},                                                 // 3 -> 4 spaces
		{"        example", "        example"},                                        // >= 4 stays the same
		{"  example", "  example"},                                                    // < 3 stays the same
		{"   example\n\texample", "    example\n\texample"},                           // multiple lines
		{"\texample", "\texample"},                                                    // tabs
		{"\texample\n   example\n    example", "\texample\n    example\n    example"}, // tabs, three spaces and four spaces
	}

	for idx, tt := range tests {
		var p Package
		if out := p.cleanDoc(tt.in); out != tt.out {
			t.Fatalf("[%v] Unexpected doc, expected=%v, got=%v", idx, tt.out, out)
		}
	}
}
