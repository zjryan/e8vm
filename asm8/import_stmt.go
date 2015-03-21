package asm8

import (
	"path"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type importStmt struct {
	*ast.ImportStmt

	as   string
	path string
	lib  *link8.Pkg
	used bool
}

func resolveImportStmt(log lex8.Logger, imp *ast.ImportStmt) *importStmt {
	ret := new(importStmt)
	ret.path = imp.Path.Lit

	if imp.As != nil {
		ret.as = imp.As.Lit
	} else {
		ret.as = path.Base(ret.path)
	}

	return ret
}
