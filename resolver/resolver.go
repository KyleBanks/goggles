// Package resolver provides the ability to locate and document go source packages.
package resolver

import (
	"log"
	"os"
	"path/filepath"

	"github.com/KyleBanks/goggles/pkg/sys"
)

// Resolver is a type that can access go packages and
// resolve their details.
type Resolver struct{}

// List returns a list of all packages in the $GOPATH.
func (r Resolver) List() ([]*Package, error) {
	ch := make(chan *Package, 0)
	expect := r.walkPackages(ch)

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
func (Resolver) Details(name string) (*Package, error) {
	p, err := NewPackage(name)
	if err != nil {
		return nil, err
	}

	if err := p.makeDocs(); err != nil {
		return nil, err
	}

	return p, nil
}

// walkPackages returns a Package on the provided channel for each package found in
// the GOPATH.
//
// The return value is the number of packages to expect to receive on the channel.
func (r Resolver) walkPackages(ch chan *Package) int {
	var count int
	visit := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		} else if !info.IsDir() {
			return nil
		}

		path = cleanPath(path)
		if ignore(path) {
			return nil
		}

		count++
		go r.resolve(path, ch)
		return nil
	}

	for _, dir := range sys.Srcdir() {
		filepath.Walk(dir, visit)
	}

	return count
}

// resolve loads the build details of a single Package and returns it
// on the provided channel.
func (Resolver) resolve(path string, ch chan *Package) {
	p, err := NewPackage(path)
	if err != nil {
		if err != errPackageNotFound {
			log.Printf("failed to resolve pkg %v: %v", path, err)
		}
		ch <- nil
		return
	}

	ch <- p
}
