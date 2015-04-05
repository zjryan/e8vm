package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type importDecl struct {
	*ast.Import

	stmts map[string]*importStmt
	paths map[string]struct{}
}

func resolveImportDecl(log lex8.Logger, imp *ast.Import) *importDecl {
	ret := new(importDecl)
	ret.Import = imp
	ret.stmts = make(map[string]*importStmt)
	ret.paths = make(map[string]struct{})

	for _, stmt := range imp.Stmts {
		r := resolveImportStmt(log, stmt)

		if other, found := ret.stmts[r.as]; found {
			log.Errorf(r.As.Pos, "%s already imported", r.as)
			log.Errorf(other.As.Pos, "  previously imported here", other.as)
			continue
		}

		ret.stmts[r.as] = r
		ret.paths[r.path] = struct{}{}
	}

	return ret
}
