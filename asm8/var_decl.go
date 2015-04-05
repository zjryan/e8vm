package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type varDecl struct {
	*ast.Var

	stmts []*varStmt
}

func resolveVar(log lex8.Logger, v *ast.Var) *varDecl {
	ret := new(varDecl)

	ret.Var = v

	for _, stmt := range v.Stmts {
		r := resolveVarStmt(log, stmt)
		ret.stmts = append(ret.stmts, r)
	}

	return ret
}
