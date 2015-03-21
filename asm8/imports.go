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

	errs := lex8.NewErrorList()
	imp := resolveImportDecl(errs, astFile.Imports)
	if len(errs.Errs) != 0 {
		return nil, errs.Errs
	}

	return imp.Paths(), nil

}
