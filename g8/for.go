package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildForStmt(b *builder, stmt *ast.ForStmt) {
	if stmt.Init == nil && stmt.Iter == nil {
		if stmt.Cond == nil {
			body := b.f.NewBlock(b.b)
			after := b.f.NewBlock(body)
			body.Jump(body)

			b.b = body

			b.breaks.push(after, "")
			b.continues.push(body, "")

			buildBlock(b, stmt.Body)

			b.breaks.pop()
			b.continues.pop()

			b.b = after
		} else {
			condBlock := b.f.NewBlock(b.b)
			body := b.f.NewBlock(condBlock)
			after := b.f.NewBlock(body)
			body.Jump(condBlock)

			b.b = condBlock
			c := buildExpr(b, stmt.Cond)
			if !c.IsBool() {
				pos := ast.ExprPos(stmt.Cond)
				b.Errorf(pos, "expect boolean expression, got %s", c)
				b.b = after
				return
			}
			condBlock.JumpIfNot(c.IR(), after)

			b.b = body

			b.breaks.push(after, "")
			b.continues.push(condBlock, "")

			buildBlock(b, stmt.Body)

			b.breaks.pop()
			b.continues.pop()

			b.b = after
		}
	} else {
		b.Errorf(stmt.Kw.Pos, "advanced for statement not implemented yet")
	}
}
