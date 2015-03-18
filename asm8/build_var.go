package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/link8"
)

func varSize(v *ast.VarDecl) int {
	ret := 0
	for _, stmt := range v.Stmts {
		ret += len(stmt.Data)
	}

	return ret
}

func varAlign(v *ast.VarDecl) uint32 {
	for _, stmt := range v.Stmts {
		if stmt.Align == 4 {
			return 4
		}
	}
	return 1
}

func buildVar(b *builder, v *ast.VarDecl) *link8.Var {
	if varSize(v) == 0 {
		b.Errorf(v.Name.Pos, "var %q has no data", v.Name.Lit)
		return nil
	}

	ret := link8.NewVar(varAlign(v))
	for _, stmt := range v.Stmts {
		n := ret.Size()
		if stmt.Align == 4 && n%4 != 0 {
			ret.Pad(4 - n%4)
		}

		ret.Write(stmt.Data)
	}

	if ret.TooLarge() {
		b.Errorf(v.Name.Pos, "var %q too large", v.Name.Lit)
		return nil
	}

	return ret
}
