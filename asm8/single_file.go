package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

// BuildSingleFile builds a package named "main" from a single file.
func BuildSingleFile(f string, rc io.ReadCloser) ([]byte, []*lex8.Error) {
	p := newParser(f, rc)
	file := parseFile(p)
	if es := p.Errs(); es != nil {
		return nil, es
	}

	pkgName := "main"
	pkg := NewPackage(pkgName)
	pkg.AddFile(file)

	b := newBuilder()
	main := buildPkg(b, pkg)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	return linkPkg(main), nil
}
