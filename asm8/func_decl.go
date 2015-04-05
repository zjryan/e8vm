package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type funcDecl struct {
	*ast.Func

	stmts []*funcStmt
}

func resolveFunc(log lex8.Logger, f *ast.Func) *funcDecl {
	ret := new(funcDecl)
	ret.Func = f

	for _, stmt := range f.Stmts {
		r := resolveFuncStmt(log, stmt)
		ret.stmts = append(ret.stmts, r)
	}

	return ret
}
