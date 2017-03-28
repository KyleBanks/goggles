package release

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var endpoint = "https://api.github.com/repos/%v/%v/releases"

// GetLatest returns the latest release name for the given repository.
func GetLatest(owner, repo string) (string, error) {
	// TODO: /latest currently returns a 404, switch when it becomes available.
	response, err := http.Get(fmt.Sprintf(endpoint, owner, repo))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var releases []struct{ Name string }
	if err := json.NewDecoder(response.Body).Decode(&releases); err != nil {
		return "", errors.Wrapf(err, "github response: %v", response.StatusCode)
	}

	if len(releases) == 0 {
		return "", nil
	}

	return releases[0].Name, nil
}

// IsLatest returns true when the provided version matches the latest release for
// the repository.
func IsLatest(owner, repo, version string) (bool, string, error) {
	v, err := GetLatest(owner, repo)
	if err != nil {
		return false, "", err
	}

	return v == version || fmt.Sprintf("v%v", version) == v, v, nil
}
