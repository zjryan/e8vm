package asm8

import (
	"sort"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type importDecl struct {
	*ast.ImportDecl

	stmts map[string]*importStmt
	paths map[string]struct{}
}

func resolveImportDecl(log lex8.Logger, imp *ast.ImportDecl) *importDecl {
	ret := new(importDecl)
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

func (imp *importDecl) Paths() []string {
	ret := make([]string, 0, len(imp.paths))
	for path := range imp.paths {
		ret = append(ret, path)
	}

	sort.Strings(ret)
	return ret
}
