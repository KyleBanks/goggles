package env

import (
	"os"
	"testing"
)

func TestEnvService_Get(t *testing.T) {
	// Defaults to dev
	os.Setenv(EnvironmentVariable, "Unknown")
	if env := Get(); env != Dev {
		t.Fatalf("Unexpected default environment: %v", env)
	}

	// Dev
	os.Setenv(EnvironmentVariable, string(Dev))
	if env := Get(); env != Dev {
		t.Fatalf("Unexpected DEV environment: %v", env)
	}

	// Test
	os.Setenv(EnvironmentVariable, string(Test))
	if env := Get(); env != Test {
		t.Fatalf("Unexpected TEST environment: %v", env)
	}

	// Prod
	os.Setenv(EnvironmentVariable, string(Prod))
	if env := Get(); env != Prod {
		t.Fatalf("Unexpected Prod environment: %v", env)
	}
}
