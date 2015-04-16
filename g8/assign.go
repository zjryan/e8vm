package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
)

func assign(b *builder, dest *ref, src *ref, op *lex8.Token) bool {
	ndest := dest.Len()
	nsrc := src.Len()
	if ndest != nsrc {
		b.Errorf(op.Pos, "cannot assign %s to %s",
			nsrc, ndest,
		)
		return false
	}

	for i, destType := range dest.typ {
		if !dest.ir[i].Addressable() {
			b.Errorf(op.Pos, "assigning to non-addressable")
			return false
		}

		srcType := src.typ[i]
		if !types.CanAssign(destType, srcType) {
			b.Errorf(op.Pos, "cannot assign %s to %s", src, dest)
			return false
		}
	}

	// perform the assignment
	for i, dest := range dest.ir {
		b.b.Assign(dest, src.ir[i])
	}

	return true
}

func buildAssignStmt(b *builder, stmt *ast.AssignStmt) {
	left := buildExprList(b, stmt.Left)
	if left == nil {
		return
	}
	right := buildExprList(b, stmt.Right)
	if right == nil {
		return
	}
	assign(b, left, right, stmt.Assign)
}
