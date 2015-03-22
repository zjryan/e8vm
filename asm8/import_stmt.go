package asm8

import (
	"path"
	"strconv"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
	"lonnie.io/e8vm/pkg8"
)

type importStmt struct {
	*ast.ImportStmt

	as   string
	path string

	linkable pkg8.Linkable
	lib      *link8.Pkg
}

func resolveImportStmt(log lex8.Logger, imp *ast.ImportStmt) *importStmt {
	ret := new(importStmt)
	ret.ImportStmt = imp

	s, e := strconv.Unquote(imp.Path.Lit)
	if e != nil {
		log.Errorf(imp.Path.Pos, "invalid string %s",
			imp.Path.Lit)
		return nil
	}

	ret.path = s

	if imp.As != nil {
		ret.as = imp.As.Lit
	} else {
		ret.as = path.Base(ret.path)
	}

	return ret
}
