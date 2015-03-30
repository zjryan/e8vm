package g8

import (
	"fmt"

	"lonnie.io/e8vm/g8/ast"
)

func buildBinaryOpExpr(b *builder, expr *ast.OpExpr) *ref {
	op := expr.Op.Lit
	A := buildExpr(b, expr.A)
	B := buildExpr(b, expr.B)
	if A == nil || B == nil { // some error occured
		return nil
	}

	opPos := expr.Op.Pos
	if bothBasic(A.typ, B.typ, typInt) {
		switch op {
		case "+", "-", "*", "&", "|":
			ret := newRef(A.typ, b.f.NewTemp(4))
			b.b.Arith(ret.ir, A.ir, op, B.ir)
			return ret
		case "%", "/":
			// TODO: division requires panic for 0
			// this would require support on if and panic
			ret := newRef(A.typ, b.f.NewTemp(4))
			b.b.Arith(ret.ir, A.ir, op, B.ir)
			return ret
		default:
			b.Errorf(opPos, "%q on ints", op)
			return nil
		}
	}

	b.Errorf(opPos, "invalid %q", op)
	return nil
}

func buildUnaryOpExpr(b *builder, expr *ast.OpExpr) *ref {
	op := expr.Op.Lit
	B := buildExpr(b, expr.B)
	if B == nil {
		return nil
	}

	opPos := expr.Op.Pos
	if isBasic(B.typ, typInt) {
		switch op {
		case "+", "-", "^":
			ret := newRef(B.typ, b.f.NewTemp(4))
			b.b.Arith(ret.ir, nil, op, B.ir)
			return ret
		default:
			b.Errorf(opPos, "%q on int", op)
			return nil
		}
	} else if isBasic(B.typ, typBool) {
		switch op {
		case "!":
			panic("todo")
		default:
			b.Errorf(opPos, "%q on boolean", op)
			return nil
		}
	}

	b.Errorf(opPos, "invalid unary operator %q", op)
	return nil
}

func buildOpExpr(b *builder, expr *ast.OpExpr) *ref {
	if expr.A == nil {
		buildUnaryOpExpr(b, expr)
	}
	return buildBinaryOpExpr(b, expr)
}

func buildCallExpr(b *builder, expr *ast.CallExpr) *ref {
	f := buildExpr(b, expr)
	_ = f
	panic("todo")
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
	case *ast.CallExpr:
		return buildCallExpr(b, expr)
	default:
		panic(fmt.Errorf("%T: invalid or not implemented", expr))
	}
}
