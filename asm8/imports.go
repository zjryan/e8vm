package asm8

import (
	"io"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

func listImport(f string, rc io.ReadCloser) ([]string, []*lex8.Error) {
	astFile, es := parse.File(f, rc)
	if es != nil {
		return nil, es
	}

	if astFile.Imports == nil {
		return nil, nil
	}

	res := newResolver()
	imp := resolveImportDecl(res, astFile.Imports)
	if es := res.Errs(); es != nil {
		return nil, es
	}

	return imp.Paths(), nil
}
