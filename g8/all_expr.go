package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildExpr(b *builder, expr ast.Expr) *ref {
	if expr == nil {
		return nil
	}

	switch expr := expr.(type) {
	case *ast.Operand:
		return buildOperand(b, expr)
	case *ast.ParenExpr:
		return buildExpr(b, expr.Expr)
	case *ast.OpExpr:
		return buildOpExpr(b, expr)
	case *ast.CallExpr:
		return buildCallExpr(b, expr)
	default:
		b.Errorf(ast.ExprPos(expr), "invalid or not implemented: %T", expr)
		return nil
	}
}
