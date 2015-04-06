package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
)

func buildType(b *builder, expr ast.Expr) types.T {
	if expr == nil {
		panic("bug")
	}

	var ret *ref
	switch expr := expr.(type) {
	case *ast.Operand:
		ret = buildOperand(b, expr)
	default:
		b.Errorf(ast.ExprPos(expr), "expect a type")
		return nil
	}

	if ret == nil {
		return nil
	} else if !ret.IsType() {
		b.Errorf(ast.ExprPos(expr), "expect a type")
		return nil
	}

	return ret.TypeType()
}

func allocTypedVars(b *builder, toks []*lex8.Token, t types.T) *ref {
	ts := make([]types.T, len(toks))
	for i := range toks {
		ts[i] = t
	}
	return allocVars(b, toks, ts)
}

func zero(b *builder, ref *ref) {
	for _, r := range ref.ir {
		b.b.Zero(r)
	}
}

func buildVarDecl(b *builder, d *ast.VarDecl) {
	idents := d.Idents.Idents

	if d.Eq != nil {
		right := buildExprList(b, d.Exprs)
		if right == nil {
			return
		}
		if d.Type != nil {
			tdest := buildType(b, d.Type)
			if tdest == nil {
				return
			}

			dest := allocTypedVars(b, idents, tdest)
			if assign(b, dest, right, d.Eq) {
				declareVars(b, idents, dest)
			}
		} else {
			define(b, idents, right, d.Eq)
		}
		return
	}

	if d.Type == nil {
		panic("must have a type")
	}

	t := buildType(b, d.Type)
	if t == nil {
		return
	}

	for _, ident := range idents {
		r := allocVar(b, ident, t)
		zero(b, r)
		declareVar(b, ident, r)
	}
}

func buildVarDecls(b *builder, decls *ast.VarDecls) {
	for _, d := range decls.Decls {
		buildVarDecl(b, d)
	}
}
