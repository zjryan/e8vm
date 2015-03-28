package g8

import (
	"fmt"
	"math"
	"strconv"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
)

func buildInt(b *builder, op *lex8.Token) *ref {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		b.Errorf(op.Pos, "invalid integer: %s", e)
		return nil
	}

	if ret < math.MinInt32 {
		b.Errorf(op.Pos, "integer too small to fit in 32-bit")
		return nil
	} else if ret > math.MaxUint32 {
		b.Errorf(op.Pos, "integer too large to fit in 32-bit")
		return nil
	} else if ret > math.MaxInt32 {
		// must be unsigned integer
		return newRef(typUint, ir.Num(uint32(ret)))
	}

	return newRef(typInt, ir.Snum(int32(ret)))
}

func buildIdent(b *builder, op *lex8.Token) *ref {
	s := b.scope.Query(op.Lit)
	if s == nil {
		b.Errorf(op.Pos, "undefined identifer %s", op.Lit)
		return nil
	}

	switch s.Type {
	case symVar:
		v := s.Item.(*objVar)
		return v.ref
	default:
		panic("todo")
	}
}

func buildOperand(b *builder, op *ast.Operand) *ref {
	switch op.Token.Type {
	case parse.Int:
		return buildInt(b, op.Token)
	case parse.Ident:
		return buildIdent(b, op.Token)
	default:
		panic("invalid or not implemented")
	}
}

func buildBinaryOpExpr(b *builder, expr *ast.OpExpr) *ref {
	op := expr.Op.Lit
	A := buildExpr(b, expr.A)
	B := buildExpr(b, expr.B)
	if A == nil || B == nil { // some error occured
		return nil
	}

	if !bothBasic(A.typ, B.typ, typInt) {
		b.Errorf(expr.Op.Pos, "we only support int operators now")
		return nil
	}

	switch op {
	case "+", "-", "*", "&", "|":
		ret := newRef(A.typ, b.f.NewTemp(4))
		b.b.Arith(ret.ir, A.ir, op, B.ir)
		return ret
	case "%", "/":
		// TODO: division requires panic for 0
		ret := newRef(A.typ, b.f.NewTemp(4))
		b.b.Arith(ret.ir, A.ir, op, B.ir)
		return ret
	default:
		panic("todo")
	}
}

func buildOpExpr(b *builder, expr *ast.OpExpr) *ref {
	if expr.A == nil {
		panic("todo: unary op")
	}
	return buildBinaryOpExpr(b, expr)
}

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
	default:
		panic(fmt.Errorf("%T: invalid or not implemented", expr))
	}
}

func buildExprList(b *builder, list *ast.ExprList) []*ref {
	ret := make([]*ref, 0, list.Len())
	for _, expr := range list.Exprs {
		ref := buildExpr(b, expr)
		if ref == nil {
			return nil
		}
		ret = append(ret, ref)
	}
	return ret
}

func buildIdentList(b *builder, list *ast.ExprList) ([]*lex8.Token, ast.Expr) {
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
