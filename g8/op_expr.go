package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
)

func buildBinaryOpExpr(b *builder, expr *ast.OpExpr) *ref {
	op := expr.Op.Lit
	A := b.buildExpr(expr.A)
	B := b.buildExpr(expr.B)
	if A == nil || B == nil { // some error occured
		return nil
	}

	opPos := expr.Op.Pos

	if !A.IsSingle() || !B.IsSingle() {
		b.Errorf(opPos, "%q on expression list", op)
		return nil
	}

	atyp := A.Type()
	btyp := B.Type()

	if types.BothBasic(atyp, btyp, types.Int) {
		switch op {
		case "+", "-", "*", "&", "|":
			ret := b.newTemp(types.Int)
			b.b.Arith(ret.IR(), A.IR(), op, B.IR())
			return ret
		case "%", "/":
			// TODO: division requires panic for 0
			// this would require support on if and panic
			ret := b.newTemp(types.Int)
			b.b.Arith(ret.IR(), A.IR(), op, B.IR())
			return ret
		case "==", "!=", ">", "<", ">=", "<=":
			ret := b.newTemp(types.Bool)
			b.b.Arith(ret.IR(), A.IR(), op, B.IR())
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
	B := b.buildExpr(expr.B)
	if B == nil {
		return nil
	}

	opPos := expr.Op.Pos

	if !B.IsSingle() {
		b.Errorf(opPos, "%q on expression list", op)
		return nil
	}

	btyp := B.Type()
	if types.IsBasic(btyp, types.Int) {
		switch op {
		case "+", "-", "^":
			ret := b.newTemp(types.Int)
			b.b.Arith(ret.IR(), nil, op, B.IR())
			return ret
		default:
			b.Errorf(opPos, "%q on int", op)
			return nil
		}
	} else if types.IsBasic(btyp, types.Bool) {
		switch op {
		case "!":
			ret := b.newTemp(types.Bool)
			b.b.Arith(ret.IR(), nil, op, B.IR())
			return ret
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
		return buildUnaryOpExpr(b, expr)
	}
	return buildBinaryOpExpr(b, expr)
}
