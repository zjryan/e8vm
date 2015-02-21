package asm8

import (
	"io"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Pkg contains the information required to build a package
type Pkg struct {
	Path    string
	Imports map[string]*ast.PkgImport
	Files   map[string]io.ReadCloser
}

// Build builds a package.
func (pb *Pkg) Build() (*link8.Pkg, []*lex8.Error) {
	pkg := ast.NewPkg(pb.Path)
	for f, rc := range pb.Files {
		parsed, es := parse.File(f, rc)
		if es != nil {
			return nil, es
		}

		pkg.AddFile(parsed)
	}

	pkg.Imports = pb.Imports

	b := newBuilder()
	ret := buildLib(b, pkg)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	return ret.Pkg, nil
}
