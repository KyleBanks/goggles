package goggles

import (
	"bytes"
	"fmt"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"

	"github.com/KyleBanks/depth"
	"github.com/KyleBanks/goggles/pkg/sys"
)

const (
	// PackageDoc indicates package-level documentation.
	PackageDoc DocType = "PACKAGE"
	// FunctionDoc indicates function-level documentation.
	FunctionDoc DocType = "FUNCTION"
	// TypeDoc indicates type-level documentation.
	TypeDoc DocType = "TYPE"
)

// DocType defines a type of documentation.
type DocType string

// Package represents a go source package.
type Package struct {
	depth.Pkg

	files *token.FileSet
	Docs  *Doc `json:"docs"`
}

// Doc represents documentation for a function, type, or package.
type Doc struct {
	Type DocType `json:"type"`

	Name        string `json:"name"`
	Header      string `json:"header"`
	Import      string `json:"import"`
	Declaration string `json:"declaration"`
	Usage       string `json:"usage"`

	Constants string `json:"constants"`
	Variables string `json:"variables"`
	Content   []Doc  `json:"content"`
}

// NewPackage attempts to resolve a go package from the path provided.
//
// The path can be either the absolute path (ex. /foo/bar/package)
// or the import path (ex. github.com/foo/bar).
func NewPackage(path string) (*Package, error) {
	t := depth.Tree{
		ResolveTest:     false,
		ResolveInternal: false,
		MaxDepth:        1,
	}
	if err := t.Resolve(path); err != nil {
		return nil, err
	}

	return &Package{
		Pkg: *t.Root,
	}, nil
}

// makeDocs retrieves the documentation for a package and attaches it to the Package.
func (p *Package) makeDocs() error {
	p.files = token.NewFileSet()
	doc, err := p.parseDocs()
	if err != nil {
		return err
	}

	p.Docs = &Doc{
		Type: PackageDoc,

		Name:   doc.Name,
		Import: fmt.Sprintf("import \"%v\"", p.Name),
		Usage:  doc.Doc,

		Constants: p.printValues(doc.Consts),
		Variables: p.printValues(doc.Vars),
	}
	p.Docs.Content = append(p.Docs.Content, p.printFuncs(doc.Funcs)...)
	p.Docs.Content = append(p.Docs.Content, p.printTypes(doc.Types)...)

	return nil
}

// parseDocs parses the package documentation.
func (p *Package) parseDocs() (*doc.Package, error) {
	filter := func(file os.FileInfo) bool {
		name := file.Name()
		return !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
	}

	pkgs, err := parser.ParseDir(p.files, sys.AbsPath(p.Name), filter, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		return doc.New(pkg, ".", 0), nil
	}

	return nil, nil
}

func (p *Package) printValues(vals []*doc.Value) string {
	var b bytes.Buffer
	for _, v := range vals {
		fmt.Fprintf(&b, "%s\n%s", p.printToken(v.Decl), p.printToken(v.Doc))
	}
	return b.String()
}

func (p *Package) printFuncs(funcs []*doc.Func) []Doc {
	var docs []Doc
	for _, f := range funcs {
		var receiver string
		if f.Recv != "" {
			receiver = fmt.Sprintf("(%s) ", f.Recv)
		}

		docs = append(docs, Doc{
			Type: FunctionDoc,

			Name:        f.Name,
			Usage:       f.Doc,
			Header:      fmt.Sprintf("func %v%v", receiver, f.Name),
			Declaration: p.printToken(f.Decl),
		})
	}

	return docs
}

func (p *Package) printTypes(types []*doc.Type) []Doc {
	var docs []Doc
	for _, t := range types {
		d := Doc{
			Type: TypeDoc,

			Name:        t.Name,
			Usage:       t.Doc,
			Header:      fmt.Sprintf("type %v", t.Name),
			Declaration: p.printToken(t.Decl),

			Constants: p.printValues(t.Consts),
			Variables: p.printValues(t.Vars),
		}
		d.Content = append(d.Content, p.printFuncs(t.Funcs)...)
		d.Content = append(d.Content, p.printFuncs(t.Methods)...)

		docs = append(docs, d)
	}

	return docs
}

func (p *Package) printToken(t interface{}) string {
	var b bytes.Buffer
	conf := printer.Config{
		Mode:     printer.UseSpaces,
		Tabwidth: 4,
	}
	err := conf.Fprint(&b, p.files, t)
	if err != nil {
		return ""
	}

	return b.String()
}
