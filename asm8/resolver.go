package asm8

import (
	"lonnie.io/e8vm/lex8"
	// "lonnie.io/e8vm/asm8/ast"
)

type resolver struct {
	*lex8.ErrorList
}

func newResolver() *resolver {
	ret := new(resolver)
	ret.ErrorList = lex8.NewErrorList()
	return ret
}

/*
func (r resolver) resolveFuncStmt(stmt *ast.FuncStmt) {
	if parseLabel(p, op0) {
		if len(ops) > 1 {
			p.Errorf(op0.Pos, "label should take the entire line")
			return nil
		}
		return &ast.FuncStmt{Label: lead, Ops: ops}
	}

	return &ast.FuncStmt{Inst: parseInst(p, ops), Ops: ops}
}

func resolveFile(f *ast.File) {
}
*/
