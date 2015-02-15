package asm8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// BuildSingleFile builds a package named "main" from a single file.
func BuildSingleFile(f string, rc io.ReadCloser) ([]byte, []*lex8.Error) {
	b := &PkgBuild{
		Path:   "main",
		Import: nil,
		Files:  map[string]io.ReadCloser{f: rc},
	}

	p, es := b.Build()
	if es != nil {
		return nil, es
	}

	ret, e := link8.LinkMain(p)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return ret, nil
}
