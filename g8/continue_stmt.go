package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildContinueStmt(b *builder, stmt *ast.ContinueStmt) {
	if stmt.Label != nil {
		b.Errorf(stmt.Label.Pos, "continue with label not implemented")
		return
	}

	next := b.continues.top()
	if next == nil {
		b.Errorf(stmt.Kw.Pos, "continue is not in a for block")
		return
	}

	after := b.f.NewBlock(b.b)
	b.b.Jump(b.continues.top())
	b.b = after
}
