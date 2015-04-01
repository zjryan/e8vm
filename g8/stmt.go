package g8

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

func allocVars(b *builder, toks []*lex8.Token, ts []types.Type) *ref {
	ret := new(ref)
	ret.typ = ts
	for i, t := range ts {
		name := toks[i].Lit      // just for debugging on IR
		v := b.newLocal(t, name) // not declared yet
		ret.ir = append(ret.ir, v)
	}
	return ret
}

func assign(b *builder, dest *ref, src *ref, op *lex8.Token) bool {
	ndest := dest.Len()
	nsrc := src.Len()
	if ndest != nsrc {
		b.Errorf(op.Pos, "cannot assign %d to %d expresssions",
			nsrc, ndest,
		)
		return false
	}

	for i, destType := range dest.typ {
		if !addressable(dest.ir[i]) {
			b.Errorf(op.Pos, "assigning to non-addressable")
			return false
		}

		srcType := src.typ[i]
		if !types.CanAssign(destType, srcType) {
			b.Errorf(op.Pos, "cannot assign %s to %s", src, dest)
			return false
		}
	}

	// perform the assignment
	for i, dest := range dest.ir {
		b.b.Assign(dest, src.ir[i])
	}

	return true
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

	idents, err := buildIdentList(b, stmt.Left)
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

func buildAssignStmt(b *builder, stmt *ast.AssignStmt) {
	left := buildExprList(b, stmt.Left)
	if left == nil {
		return
	}
	right := buildExprList(b, stmt.Right)
	if right == nil {
		return
	}
	assign(b, left, right, stmt.Assign)
}

func buildStmt(b *builder, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		buildExpr(b, stmt.Expr)
		// TODO: check if expr is a function call
	case *ast.DefineStmt:
		buildDefineStmt(b, stmt)
	case *ast.AssignStmt:
		buildAssignStmt(b, stmt)
	default:
		panic("invalid or not implemented")
	}
}
