package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
)

func buildExprList(b *builder, list *ast.ExprList) *ref {
	if list == nil {
		return new(ref) // empty ref, for void
	}

	n := list.Len()

	ret := new(ref)
	if n == 0 {
		return ret // empty ref
	} else if n == 1 {
		for _, expr := range list.Exprs {
			return b.buildExpr(expr)
		}
		panic("unreachable")
	}

	for _, expr := range list.Exprs {
		ref := b.buildExpr(expr)
		if ref == nil {
			return nil
		}
		if !ref.IsSingle() {
			b.Errorf(ast.ExprPos(expr), "cannot composite list in a list")
			return nil
		}

		ret.typ = append(ret.typ, ref.Type())
		ret.ir = append(ret.ir, ref.IR())
	}

	return ret
}

func buildIdentExprList(b *builder, list *ast.ExprList) (
	idents []*lex8.Token, firstError ast.Expr,
) {
	ret := make([]*lex8.Token, 0, list.Len())
	for _, expr := range list.Exprs {
		op, ok := expr.(*ast.Operand)
		if !ok {
			return nil, expr
		}
		if op.Token.Type != parse.Ident {
			return nil, expr
		}

		ret = append(ret, op.Token)
	}

	return ret, nil
}
