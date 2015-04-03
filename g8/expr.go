package g8

import (
	"fmt"

	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
)

func buildBinaryOpExpr(b *builder, expr *ast.OpExpr) *ref {
	op := expr.Op.Lit
	A := buildExpr(b, expr.A)
	B := buildExpr(b, expr.B)
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
	B := buildExpr(b, expr.B)
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

func buildCallExpr(b *builder, expr *ast.CallExpr) *ref {
	f := buildExpr(b, expr.Func)
	if f == nil {
		return nil
	}

	pos := ast.ExprPos(expr.Func)

	if !f.IsSingle() {
		b.Errorf(pos, "expression list is not callable")
		return nil
	}

	funcType, ok := f.Type().(*types.Func) // the func sig in the builder
	if !ok {
		// not a function
		b.Errorf(pos, "function call on non-callable")
		return nil
	}

	args := buildExprList(b, expr.Args)
	if args == nil {
		return nil
	}

	if args.Len() != len(funcType.Args) {
		b.Errorf(ast.ExprPos(expr), "argument expects (%s), got (%s)",
			fmt8.Join(funcType.Args, ","), fmt8.Join(args.typ, ","),
		)
		return nil
	}

	// type check on parameters
	for i, argType := range args.typ {
		expect := funcType.Args[i].Type
		if !types.CanAssign(expect, argType) {
			pos := ast.ExprPos(expr.Args.Exprs[i])
			b.Errorf(pos, "argument %d expects %s, got %s",
				i, expect, argType,
			)
		}
	}

	ret := new(ref)
	ret.typ = funcType.RetTypes
	for _, t := range funcType.RetTypes {
		ret.ir = append(ret.ir, b.f.NewTemp(t.Size()))
	}

	// call the func in IR
	b.b.Call(ret.ir, f.IR(), funcType.Sig, args.ir...)

	return ret
}

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
			return buildExpr(b, expr)
		}
		panic("unreachable")
	}

	for _, expr := range list.Exprs {
		ref := buildExpr(b, expr)
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

func buildIdentList(b *builder, list *ast.ExprList) (
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
		b.Errorf(ast.ExprPos(expr), "%s",
			fmt.Errorf("%T: invalid or not implemented", expr),
		)
		return nil
	}
}
