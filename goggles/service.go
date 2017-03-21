package goggles

import (
	"log"
	"os"
	"path/filepath"

	"github.com/KyleBanks/depth"
	"github.com/KyleBanks/goggles/pkg/sys"
)

// Default is the default Service.
var Default Service

// Service is a pkg service that can access go packages.
type Service struct{}

// List returns a list of all packages in the $GOPATH.
func (s Service) List() ([]*Package, error) {
	ch := make(chan *Package, 0)
	expect := s.walkPackages(ch)

	// Wait for all the results and append them to the slice.
	var pkgs []*Package
	for i := 0; i < expect; i++ {
		p := <-ch
		if p != nil {
			pkgs = append(pkgs, p)
		}
	}

	return pkgs, nil
}

// Details returns the full details of a Package.
func (Service) Details(name string) (*Package, error) {
	p, err := NewPackage(name)
	if err != nil {
		return nil, err
	}

	if err := p.makeDocs(); err != nil {
		return nil, err
	}

	return p, nil
}

// walkPackages sends a Package on the provided channel for each package found in
// the gopath.
//
// The return value is the number of packages to expect to receive on the channel.
func (Service) walkPackages(ch chan *Package) int {
	var count int
	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		} else if !info.IsDir() {
			return nil
		}

		path = cleanPath(path)
		if ignorePkg(path) {
			return nil
		}

		count++
		go func(path string, ch chan *Package) {
			p, err := NewPackage(path)
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
	}

	filepath.Walk(sys.Srcdir(), visit)
	return count
}
