package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

func allocVar(b *builder, tok *lex8.Token, t types.T) *ref {
	return newRef(t, b.newLocal(t, tok.Lit))
}

func allocVars(b *builder, toks []*lex8.Token, ts []types.T) *ref {
	ret := new(ref)
	ret.typ = ts
	for i, tok := range toks {
		// the name here is just for debugging
		// it is not a var declare
		v := b.newLocal(ts[i], tok.Lit)
		ret.ir = append(ret.ir, v)
	}
	return ret
}

func declareVar(b *builder, tok *lex8.Token, r *ref) {
	name := tok.Lit
	v := &objVar{name, r}
	s := sym8.Make(name, symVar, v, tok.Pos)
	conflict := b.scope.Declare(s)
	if conflict != nil {
		b.Errorf(tok.Pos, "%q already declared as a %s",
			name, symStr(conflict.Type),
		)
	}
}

func declareVars(b *builder, toks []*lex8.Token, r *ref) {
	for i, t := range r.typ {
		declareVar(b, toks[i], newRef(t, r.ir[i]))
	}
}

func define(b *builder, idents []*lex8.Token, expr *ref, eq *lex8.Token) {
	// check count matching
	nleft := len(idents)
	nright := expr.Len()
	if nleft != nright {
		b.Errorf(eq.Pos,
			"defined %d identifers with %d expressions",
			nleft, nright,
		)
		return
	}

	left := allocVars(b, idents, expr.typ)
	if assign(b, left, expr, eq) {
		declareVars(b, idents, left)
	}
}

func buildDefineStmt(b *builder, stmt *ast.DefineStmt) {
	right := buildExprList(b, stmt.Right)
	if right == nil { // an error occured on the expression list
		return
	}

	idents, err := buildIdentExprList(b, stmt.Left)
	if err != nil {
		b.Errorf(ast.ExprPos(err), "left side of := must be identifer")
		return
	}

	define(b, idents, right, stmt.Define)
}
