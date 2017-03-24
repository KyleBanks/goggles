// Package storage provides the ability to persist and retrieve structs.
package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// FileStore represents a simple on-disk file storage system.
type FileStore struct {
	Dirname  string
	Filename string
}

// NewFileStore returns an initialized FileStore.
func NewFileStore(d, f string) *FileStore {
	return &FileStore{
		Dirname:  d,
		Filename: f,
	}
}

// Load attempts to read and decode the storage file into the
// value provided.
func (f FileStore) Load(v interface{}) error {
	if err := f.makePath(); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(f.path())
	if err != nil {
		return err
	}

	json.Unmarshal(data, v)
	return nil
}

// Save attempts to write the value provided out to the storage file.
func (f FileStore) Save(v interface{}) error {
	if err := f.makePath(); err != nil {
		return err
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(f.path(), data, os.ModePerm)
}

// makePath attempts to create the neccessary directories to the Filestore
// cache file location.
func (f FileStore) makePath() error {
	return os.MkdirAll(filepath.Dir(f.path()), os.ModePerm)
}

// CacheDir returns the cache path for the current platform.
func (f FileStore) path() string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches", f.Dirname, f.Filename)
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), f.Dirname, f.Filename)
	default:
		return filepath.Join(os.Getenv("HOME"), ".cache", f.Dirname, f.Filename)
	}
}
