package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
)

func buildOperand(b *builder, op *ast.Operand) ir.Ref {
	switch op.Token.Type {
	case parse.Int:
		panic("todo: integer")
	default:
		panic("invalid or not implemented")
	}
}

func buildExpr(b *builder, expr ast.Expr) ir.Ref {
	if expr == nil {
		return nil
	}

	switch expr := expr.(type) {
	default:
		_ = expr
		panic("invalid or not implemented")
	}
}

func buildExprList(b *builder, list *ast.ExprList) []ir.Ref {
	ret := make([]ir.Ref, 0, list.Len())
	for _, expr := range list.Exprs {
		ref := buildExpr(b, expr)
		if ref == nil {
			return nil
		}
		ret = append(ret, ref)
	}
	return ret
}
