package gonamo

import (
	"testing"
)

func TestDefaultOptions(t *testing.T) {
	opts := defaultOptions()
	if opts == nil {
		t.Fatal("Expected defaultOptions to not return nil.")
	}

	if len(opts.Endpoint) > 0 {
		t.Fatalf("Expected Endpoint to be empty, got=%v", opts.Endpoint)
	}
	if len(opts.Region) > 0 {
		t.Fatalf("Expected Region to be empty, got=%v", opts.Region)
	}
	if opts.DefaultProvisioning <= 0 {
		t.Fatalf("Expected DefaultProvisioning to be > 0, got=%v", opts.DefaultProvisioning)
	}
}
