package goggles

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/KyleBanks/depth"
)

var (
	ignorePaths = []string{".git", "/vendor/", "/testdata/"}
)

func init() {
	log.Printf("$GOPATH=%v, srcdir=%v", gopath(), srcdir())
}

// ListPkgs returns a list of all packages in the $GOPATH.
func ListPkgs() ([]*Pkg, error) {
	ch := make(chan *Pkg, 0)
	var expect int

	filepath.Walk(srcdir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		} else if !info.IsDir() {
			return nil
		}

		path = cleanPath(path)
		if ignorePkg(path) {
			return nil
		}

		expect++
		go func(path string, ch chan *Pkg) {
			p, err := resolve(path)
			if err != nil {
				if err != depth.ErrRootPkgNotResolved {
					log.Printf("failed to resolve pkg %v: %v", path, err)
				}
				ch <- nil
				return
			}

			ch <- p
		}(path, ch)
		return nil
	})

	// Wait for all the results and append them to the slice.
	var pkgs []*Pkg
	for i := 0; i < expect; i++ {
		p := <-ch
		if p != nil {
			pkgs = append(pkgs, p)
		}
	}

	return pkgs, nil
}

// Details returns the full details of a Pkg.
func Details(name string) (*Pkg, error) {
	p, err := resolve(name)
	if err != nil {
		return nil, err
	}

	if err := p.makeDocs(); err != nil {
		return nil, err
	}

	return p, nil
}

// resolve attempts to resolve a single package and passes it to the channel provided.
func resolve(path string) (*Pkg, error) {
	t := depth.Tree{
		ResolveTest:     false,
		ResolveInternal: false,
		MaxDepth:        1,
	}
	if err := t.Resolve(path); err != nil {
		if err != depth.ErrRootPkgNotResolved {
			log.Printf("failed to resolve pkg %v: %v", path, err)
		}
		return nil, err
	}

	return &Pkg{
		Pkg: *t.Root,
	}, nil
}

// cleanPath sanitizes a package path.
func cleanPath(path string) string {
	path = strings.Replace(path, srcdir(), "", 1)
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	return path
}

// ignorePkg checks if the path provided should be ignored.
func ignorePkg(path string) bool {
	if len(path) == 0 {
		return true
	}

	for _, s := range ignorePaths {
		if strings.Contains(path, s) {
			return true
		}
	}

	return false
}

// srcdir returns the source directory for go packages.
func srcdir() string {
	return filepath.Join(gopath(), "src")
}

// gopath returns the $GOPATH environment variable.
func gopath() string {
	return os.Getenv("GOPATH")
}
