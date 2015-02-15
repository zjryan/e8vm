package asm8

import (
	"os"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildPkg builds a package from a list of files.
func BuildPkg(path string, files []string) ([]byte, []*lex8.Error) {
	pkg := newPkg(path)
	for _, f := range files {
		reader, e := os.Open(f)
		if e != nil {
			return nil, lex8.SingleErr(e)
		}

		p := newParser(f, reader)
		parsed := parseFile(p)
		if es := p.Errs(); es != nil {
			return nil, es
		}

		pkg.AddFile(parsed)
	}

	b := newBuilder()
	lib := buildLib(b, pkg)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	ret, e := link8.LinkMain(lib.Package)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return ret, nil
}
