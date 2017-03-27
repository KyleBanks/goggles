package release

import (
	"testing"
)

func Test_GetLatest(t *testing.T) {
	tests := []struct {
		owner         string
		repo          string
		expectSuccess bool
	}{
		{"KyleBanks", "goggles", true},
		{"KyleBanks", "definitely doesnt exist", false},
	}

	for idx, tt := range tests {
		res, err := GetLatest(tt.owner, tt.repo)
		if (err != nil || len(res) == 0) && tt.expectSuccess {
			t.Fatalf("[%v] Expected response, got err=%v, res=%v", idx, err, res)
		} else if !tt.expectSuccess && err == nil {
			t.Fatalf("[%v] Expected err, got nil", idx)
		}
	}
}

func Test_IsLatest(t *testing.T) {
	// error case
	{
		_, _, err := IsLatest("KyleBanks", "not a repo", "123")
		if err == nil {
			t.Fatal("Expect err, got nil")
		}
	}

	latest, err := GetLatest("KyleBanks", "goggles")
	if err != nil {
		t.Fatal(err)
	}

	// positive case
	{
		isLatest, version, err := IsLatest("KyleBanks", "goggles", latest)
		if err != nil {
			t.Fatal(err)
		}

		if !isLatest {
			t.Fatalf("Expected isLatest to be true, got version=%v", version)
		}
	}

	// negative case
	{
		isLatest, version, err := IsLatest("KyleBanks", "goggles", latest+".789")
		if err != nil {
			t.Fatal(err)
		}

		if isLatest {
			t.Fatalf("Expected isLatest to be false, got version=%v", version)
		}
	}
}
