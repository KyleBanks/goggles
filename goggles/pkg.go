package goggles

import (
	"bytes"
	"fmt"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/KyleBanks/depth"
)

const (
	// Package indicates package-level documentation.
	Package DocType = "PACKAGE"
	// Function indicates function-level documentation.
	Function DocType = "FUNCTION"
	// Type indicates type-level documentation.
	Type DocType = "TYPE"
)

// DocType defines a type of documentation.
type DocType string

// Pkg represents a go source package.
type Pkg struct {
	depth.Pkg

	files *token.FileSet
	Docs  Doc `json:"docs"`
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

// makeDocs retrieves the documentation for a package and attaches it to the Pkg.
func (p *Pkg) makeDocs() error {
	p.files = token.NewFileSet()
	doc, err := p.parseDocs()
	if err != nil {
		return err
	}

	p.Docs = Doc{
		Type: Package,

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
func (p *Pkg) parseDocs() (*doc.Package, error) {
	filter := func(file os.FileInfo) bool {
		name := file.Name()
		return !strings.HasPrefix(name, ".") && strings.HasSuffix(name, ".go") && !strings.HasSuffix(name, "_test.go")
	}

	pkgs, err := parser.ParseDir(p.files, filepath.Join(srcdir(), p.Name), filter, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	for _, pkg := range pkgs {
		return doc.New(pkg, ".", 0), nil
	}

	return nil, nil
}

func (p *Pkg) printValues(vals []*doc.Value) string {
	var b bytes.Buffer
	for _, v := range vals {
		fmt.Fprintf(&b, "%s\n%s", p.printToken(v.Decl), p.printToken(v.Doc))
	}
	return b.String()
}

func (p *Pkg) printFuncs(funcs []*doc.Func) []Doc {
	var docs []Doc
	for _, f := range funcs {
		var receiver string
		if f.Recv != "" {
			receiver = fmt.Sprintf("(%s) ", f.Recv)
		}

		docs = append(docs, Doc{
			Type: Function,

			Name:        f.Name,
			Usage:       f.Doc,
			Header:      fmt.Sprintf("func %v%v", receiver, f.Name),
			Declaration: p.printToken(f.Decl),
		})
	}

	return docs
}

func (p *Pkg) printTypes(types []*doc.Type) []Doc {
	var docs []Doc
	for _, t := range types {
		d := Doc{
			Type: Type,

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

func (p *Pkg) printToken(t interface{}) string {
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
