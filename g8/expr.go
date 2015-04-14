package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func buildExprStmt(b *builder, expr ast.Expr) {
	if e, ok := expr.(*ast.CallExpr); ok {
		buildCallExpr(b, e)
	} else {
		b.Errorf(ast.ExprPos(expr), "invalid expression statement")
	}
}
