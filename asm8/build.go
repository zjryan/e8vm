package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

// PkgBuild contains the information to build
// a package
type PkgBuild struct {
	Path string
	Files map[string]io.ReadCloser
	Import map[string]*Lib
}

// Build builds a package.
func (pb *PkgBuild) Build() (*Lib, []*lex8.Error) {
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

	return ret, nil
}
