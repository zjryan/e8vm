package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildIf(b *builder, cond ast.Expr, ifs ast.Stmt, elses *ast.ElseStmt) {
	c := buildExpr(b, cond)
	if !c.IsBool() {
		pos := ast.ExprPos(cond)
		b.Errorf(pos, "expect boolean expression, got %s", c)
		return
	}

	if elses == nil {
		body := b.f.NewBlock(b.b)
		after := b.f.NewBlock(body)
		b.b.JumpIfNot(c.IR(), after)
		b.b = body
		switch ifs := ifs.(type) {
		case *ast.Block:
			buildBlock(b, ifs)
		case *ast.ReturnStmt:
			buildReturnStmt(b, ifs)
		case *ast.BreakStmt:
			buildBreakStmt(b, ifs)
		case *ast.ContinueStmt:
			buildContinueStmt(b, ifs)
		default:
			b.Errorf(ast.ExprPos(cond), "short if statement not implemented")
		}
		b.b = after
		return
	}

	ifBody := b.f.NewBlock(b.b)
	elseBody := b.f.NewBlock(ifBody)
	after := b.f.NewBlock(elseBody)
	b.b.JumpIfNot(c.IR(), elseBody)
	ifBody.Jump(after)

	b.b = ifBody // switch to if body
	buildBlock(b, ifs.(*ast.Block))
	b.b = elseBody
	buildElseStmt(b, elses)

	b.b = after
}

func buildElseStmt(b *builder, stmt *ast.ElseStmt) {
	if stmt.If == nil {
		if stmt.Expr != nil {
			panic("invalid expression in else")
		}
		buildBlock(b, stmt.Body)
	} else {
		buildIf(b, stmt.Expr, stmt.Body, stmt.Next)
	}
}

func buildIfStmt(b *builder, stmt *ast.IfStmt) {
	buildIf(b, stmt.Expr, stmt.Body, stmt.Else)
}
