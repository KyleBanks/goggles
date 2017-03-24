package resolver

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
	"github.com/KyleBanks/goggles/pkg/sys"
)

const (
	// PackageDoc indicates package-level documentation.
	PackageDoc DocType = "PACKAGE"
	// FunctionDoc indicates function-level documentation.
	FunctionDoc DocType = "FUNCTION"
	// TypeDoc indicates type-level documentation.
	TypeDoc DocType = "TYPE"

	travisFile = ".travis.yml"
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
	Repository  string `json:"repository"`
	Header      string `json:"header"`
	Import      string `json:"import"`
	Declaration string `json:"declaration"`
	Usage       string `json:"usage"`

	Constants string `json:"constants"`
	Variables string `json:"variables"`
	Content   []Doc  `json:"content"`

	HasTravis bool `json:"hasTravis"`
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

	// Need to Chdir into the package directory.
	//
	// This only applies because the Goggles application itself is a package,
	// and if you have Goggles and one of its dependencies in your GOPATH,
	// the resolver will assume you want to import the dependency from goggles/vendor/.
	os.Chdir(sys.AbsPath(path))

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

		Name:       doc.Name,
		Repository: repo(p.Name),
		Import:     fmt.Sprintf("import \"%v\"", p.Name),
		Usage:      p.cleanDoc(doc.Doc),

		Constants: p.printValues(doc.Consts),
		Variables: p.printValues(doc.Vars),
		HasTravis: p.hasTravis(),
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
		return doc.New(pkg, sys.AbsPath(p.Name), 0), nil
	}

	return nil, nil
}

func (p *Package) printValues(vals []*doc.Value) string {
	if vals == nil {
		return ""
	}

	var b bytes.Buffer
	for _, v := range vals {
		fmt.Fprintf(&b, "%s\n%s", p.printToken(v.Decl), p.printToken(v.Doc))
	}
	return b.String()
}

func (p *Package) printFuncs(funcs []*doc.Func) []Doc {
	var docs []Doc
	if funcs == nil {
		return docs
	}

	for _, f := range funcs {
		var receiver string
		if f.Recv != "" {
			receiver = fmt.Sprintf("(%s) ", f.Recv)
		}

		docs = append(docs, Doc{
			Type: FunctionDoc,

			Name:        f.Name,
			Usage:       p.cleanDoc(f.Doc),
			Header:      fmt.Sprintf("func %v%v", receiver, f.Name),
			Declaration: p.printToken(f.Decl),
		})
	}

	return docs
}

func (p *Package) printTypes(types []*doc.Type) []Doc {
	var docs []Doc
	if types == nil {
		return docs
	}

	for _, t := range types {
		d := Doc{
			Type: TypeDoc,

			Name:        t.Name,
			Usage:       p.cleanDoc(t.Doc),
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

func (p *Package) cleanDoc(doc string) string {
	lines := strings.Split(doc, "\n")
	for i, line := range lines {

		// Be a little more lenient on the code blocks, allow three spaces
		// instead of requiring four.
		//replace(/\n   +/g, '\n\t')
		if strings.HasPrefix(line, "   ") && !strings.HasPrefix(line, "    ") {
			lines[i] = " " + line
		}
	}

	return strings.Join(lines, "\n")
}

// hasTravis returns true if the current Package or the root directory of the repository
// has a .travis.yml file present.
func (p *Package) hasTravis() bool {
	paths := []string{
		p.Name,
		repo(p.Name),
	}

	for _, p := range paths {
		_, err := os.Stat(filepath.Join(sys.AbsPath(p), travisFile))
		if err == nil {
			return true
		}
	}

	return false
}
