package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

type varDecl struct {
	*ast.VarDecl

	stmts []*varStmt
}

func resolveVar(log lex8.Logger, v *ast.VarDecl) *varDecl {
	ret := new(varDecl)

	ret.VarDecl = v

	for _, stmt := range v.Stmts {
		r := resolveVarStmt(log, stmt)
		ret.stmts = append(ret.stmts, r)
	}

	return ret
}
