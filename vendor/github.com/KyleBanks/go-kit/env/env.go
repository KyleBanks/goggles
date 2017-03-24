// Package env provides application environment detection, and support for a Dev/Test/Prod environment system.
package env

import (
	"os"
)

const (
	// EnvironmentVariable is the name of the environment variable to look for
	// when determining the application environment.
	EnvironmentVariable = "GO_ENV"
)

var (
	// Dev is a development environment.
	Dev Environment = "DEV"
	// Test is a testing environment.
	Test Environment = "TEST"
	// Prod is a production environment.
	Prod Environment = "PROD"
)

// Environment defines the name of an environment, such as Prod or Dev.
type Environment string

// Get returns the current environment that the go application is running in,
// based on the environment variable. If no environment variable is found, or
// it is not one of Dev/Test/Prod, the default (Dev) will be returned.
func Get() Environment {
	envVar := os.Getenv(EnvironmentVariable)

	switch Environment(envVar) {
	case Test:
		return Test
	case Prod:
		return Prod
	default:
		return Dev
	}
}
