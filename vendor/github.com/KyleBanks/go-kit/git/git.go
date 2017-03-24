// Package git provides git source control functionality.
package git

import (
	"fmt"
	"os"
	"strings"
)

const (
	filePermission = 0755
)

// InstallPreCommitHook takes a string to set as the pre-commit git hook
// for the .git directory specified by the path provided.
//
// The path should be to the parent of the .git directory.
// For example: <gopath>/src/github.com/KyleBanks/go-kit
func InstallPreCommitHook(hook string, path string) error {
	if len(path) > 0 && !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	filename := fmt.Sprintf("%v.git/hooks/pre-commit", path)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Ensure the file permissions are updated
	if err := os.Chmod(filename, filePermission); err != nil {
		return err
	}

	if _, err := file.WriteString(hook); err != nil {
		return err
	}

	return nil
}
