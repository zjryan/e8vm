package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildBreakStmt(b *builder, stmt *ast.BreakStmt) {
	if stmt.Label != nil {
		b.Errorf(stmt.Label.Pos, "break with label not implemented")
		return
	}

	after := b.f.NewBlock(b.b)
	b.b.Jump(b.breaks.top())
	b.b = after
}
