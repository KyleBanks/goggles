// Package clipboard provides the ability to read and write to the system clipboard.
//
// Note: Currently only supports Mac OS.
package clipboard

import (
	"bytes"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

const (
	cmdRead  = "pbpaste"
	cmdWrite = "pbcopy"
)

// Read returns the current contents of the system clipboard.
func Read() (io.Reader, error) {
	cmd := exec.Command(cmdRead)

	b, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(b), nil
}

// ReadString returns the current contents of the system clipboard as a string.
func ReadString() (string, error) {
	r, err := Read()
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Write stores the contents of the reader provided in the system clipboard.
func Write(r io.Reader) error {
	cmd := exec.Command(cmdWrite)
	p, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}
	if _, err := io.Copy(p, r); err != nil {
		return err
	}
	if err := p.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

// WriteString stores a string in the system clipboard.
func WriteString(s string) error {
	return Write(strings.NewReader(s))
}
