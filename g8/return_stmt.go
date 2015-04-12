package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
)

func buildFuncExit(b *builder) {
	next := b.f.NewBlock(b.b)
	b.b.Jump(b.f.End())
	b.b = next
}

func buildReturnStmt(b *builder, stmt *ast.ReturnStmt) {
	pos := stmt.Kw.Pos
	if stmt.Exprs == nil {
		println(b.fretNamed)
		if b.fretRef == nil || b.fretNamed {
			buildFuncExit(b)
		} else {
			b.Errorf(pos, "expects return %s", b.fretRef)
		}
	} else {
		if b.fretRef == nil {
			b.Errorf(pos, "function expects no return value")
			return
		}

		ref := buildExprList(b, stmt.Exprs)
		if ref == nil {
			return
		}

		nret := b.fretRef.Len()
		nsrc := ref.Len()
		if nret != nsrc {
			b.Errorf(pos, "expects (%s), returning (%s)", b.fretRef, ref)
			return
		}

		for i, t := range b.fretRef.typ {
			if !addressable(b.fretRef.ir[i]) {
				panic("bug")
			}

			srcType := ref.typ[i]
			if !types.CanAssign(t, srcType) {
				b.Errorf(pos, "expect (%s), returning (%s)", b.fretRef, ref)
				return
			}
		}

		for i, dest := range b.fretRef.ir {
			b.b.Assign(dest, ref.ir[i])
		}

		buildFuncExit(b)
	}
}
