package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNewFileStore(t *testing.T) {
	d := "dir"
	f := "file"

	fs := NewFileStore(d, f)
	if fs.Dirname != d {
		t.Fatalf("Unexpected Dirname, expected=%v, got=%v", d, fs.Dirname)
	}
	if fs.Filename != f {
		t.Fatalf("Unexpected Filename, expected=%v, got=%v", f, fs.Filename)
	}
}

func TestFileStore_LoadAndSave(t *testing.T) {
	sample := struct {
		Test time.Time
	}{
		Test: time.Now(),
	}

	d := "commuter-test"
	f := fmt.Sprintf("test-%v", time.Now().UnixNano())
	fs := NewFileStore(d, f)

	// Non-existent file
	// Expect err
	if err := fs.Load(&sample); err == nil {
		t.Fatal("Expected err, got nil")
	}

	// Save it
	if err := fs.Save(&sample); err != nil {
		t.Fatal(err)
	}

	// File exists, should work now
	expect := sample.Test
	sample.Test = time.Time{}
	if err := fs.Load(&sample); err != nil {
		t.Fatal(err)
	}

	if sample.Test.Unix() != expect.Unix() {
		t.Fatalf("Unexpected result, expected=%v, got=%v", expect, sample.Test)
	}
}

func TestFileStore_path(t *testing.T) {
	base := "/fake/path"
	os.Setenv("HOME", base)
	os.Setenv("LOCALAPPDATA", base)

	d := "dir"
	f := "file"
	fs := NewFileStore(d, f)

	path := fs.path()
	if !strings.HasPrefix(path, base) {
		t.Fatalf("Unexpected path prefix, expected=%v, got=%v", base, path)
	}
	if !strings.HasSuffix(path, filepath.Join(d, f)) {
		t.Fatalf("Unexpected path suffix, expected=%v, got=%v", filepath.Join(d, f), path)
	}
}
