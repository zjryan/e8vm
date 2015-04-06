package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

func allocVars(b *builder, toks []*lex8.Token, ts []types.T) *ref {
	ret := new(ref)
	ret.typ = ts
	for i, t := range ts {
		name := toks[i].Lit      // just for debugging on IR
		v := b.newLocal(t, name) // not declared yet
		ret.ir = append(ret.ir, v)
	}
	return ret
}

func declareVars(b *builder, toks []*lex8.Token, r *ref) {
	for i, t := range r.typ {
		tok := toks[i]
		name := tok.Lit
		v := &objVar{name, newRef(t, r.ir[i])}
		s := sym8.Make(name, symVar, v, tok.Pos)
		conflict := b.scope.Declare(s)
		if conflict != nil {
			b.Errorf(tok.Pos, "%q already declared as a %s",
				name, symStr(conflict.Type),
			)
		}
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

	// check count matching
	nleft := len(idents)
	nright := right.Len()
	if nleft != nright {
		b.Errorf(stmt.Define.Pos,
			"defined %d identifers with %d expressions",
			nleft, nright,
		)
		return
	}

	left := allocVars(b, idents, right.typ)
	if assign(b, left, right, stmt.Define) {
		declareVars(b, idents, left)
	}
}
