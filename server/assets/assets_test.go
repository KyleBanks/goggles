package assets

import (
	"strings"
	"testing"
)

func Test_FS(t *testing.T) {
	fs := FS()

	data, err := fs.Asset("static/index.html")
	if err != nil {
		t.Fatal(err)
	}

	index := string(data)
	if !strings.HasPrefix(index, "<!DOCTYPE html>") || !strings.HasSuffix(index, "</html>\n") {
		t.Fatalf("Unexpected index.html: %v", index)
	}
}
