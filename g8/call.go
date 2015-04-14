package g8

import (
	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
)

func buildCallExpr(b *builder, expr *ast.CallExpr) *ref {
	f := b.buildExpr(expr.Func)
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
		expect := funcType.Args[i].T
		if !types.CanAssign(expect, argType) {
			pos := ast.ExprPos(expr.Args.Exprs[i])
			b.Errorf(pos, "argument %d expects %s, got %s",
				i+1, expect, argType,
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
