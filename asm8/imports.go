package asm8

import (
	"io"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
)

func listImport(
	f string, rc io.ReadCloser, imp build8.Importer,
) []*lex8.Error {
	astFile, es := parse.File(f, rc)
	if es != nil {
		return es
	}

	if astFile.Imports == nil {
		return nil
	}

	log := lex8.NewErrorList()
	impDecl := resolveImportDecl(log, astFile.Imports)
	if es := log.Errs(); es != nil {
		return es
	}

	for as, stmt := range impDecl.stmts {
		imp.Import(as, stmt.path, stmt.Path.Pos)
	}

	return nil
}
