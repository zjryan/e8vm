package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// PkgCompile contains the information to build a package
type PkgBuild struct {
	Path   string
	Import map[string]*lib
	Files  map[string]io.ReadCloser
}

// Build builds a package.
func (pb *PkgBuild) Build() (*link8.Package, []*lex8.Error) {
	pkg := newPkg(pb.Path)
	for f, rc := range pb.Files {
		p := newParser(f, rc)
		parsed := parseFile(p)
		if es := p.Errs(); es != nil {
			return nil, es
		}

		pkg.AddFile(parsed)
	}

	b := newBuilder()
	ret := buildLib(b, pkg)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	return ret.Package, nil
}
