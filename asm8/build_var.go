package asm8

import (
	"lonnie.io/e8vm/link8"
)

func varSize(v *varDecl) int {
	ret := 0
	for _, stmt := range v.stmts {
		ret += len(stmt.data)
	}

	return ret
}

func varAlign(v *varDecl) uint32 {
	for _, stmt := range v.stmts {
		if stmt.align == 4 {
			return 4
		}
	}
	return 1
}

func buildVar(b *builder, v *varDecl) *link8.Var {
	if varSize(v) == 0 {
		b.Errorf(v.Name.Pos, "var %q has no data", v.Name.Lit)
		return nil
	}

	ret := link8.NewVar(varAlign(v))
	for _, stmt := range v.stmts {
		n := ret.Size()
		if stmt.align == 4 && n%4 != 0 {
			ret.Pad(4 - n%4)
		}

		ret.Write(stmt.data)
	}

	if ret.TooLarge() {
		b.Errorf(v.Name.Pos, "var %q too large", v.Name.Lit)
		return nil
	}

	return ret
}
