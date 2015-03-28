package g8

import (
	"lonnie.io/e8vm/g8/ast"
)

func declareVar(b *builder, name string, t typ) *ref {
	panic("todo")
}

func buildDefineStmt(b *builder, stmt *ast.DefineStmt) {
	if stmt.Left.Len() == stmt.Right.Len() {
		rights := buildExprList(b, stmt.Right)
		if rights == nil {
			return
		}

		lefts, err := buildIdentList(b, stmt.Left)
		if err != nil {
			b.Errorf(ast.ExprPos(err), "non name on left side of :=")
			return
		}

		leftRefs := make([]*ref, 0, len(lefts))
		for i, left := range lefts {
			// TODO: ignore "_"
			r := declareVar(b, left.Lit, rights[i].typ)
			leftRefs = append(leftRefs, r)
		}

		for i, left := range leftRefs {
			b.b.Assign(left.ir, rights[i].ir)
		}
	} else if stmt.Right.Len() == 1 {
		panic("todo: right might be a func call that retunrs a list")
	} else {
		b.Errorf(stmt.Define.Pos, "mismatch definition")
	}
}

func buildStmt(b *builder, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		buildExpr(b, stmt.Expr)
		// TODO: check if expr is a function call
	case *ast.DefineStmt:
		buildDefineStmt(b, stmt)
	case *ast.AssignStmt:
		if stmt.Left.Len() == stmt.Right.Len() {
			lefts := buildExprList(b, stmt.Left)
			if lefts == nil {
				return
			}

			rights := buildExprList(b, stmt.Right)
			if rights == nil {
				return
			}

			// TODO: check type matching
			for i, left := range lefts {
				b.b.Assign(left.ir, rights[i].ir)
			}
		} else if stmt.Right.Len() == 1 {
			panic("todo: right might be a function call that retunrs a list")
		} else {
			b.Errorf(stmt.Assign.Pos, "mismatch assignment")
		}
	default:
		panic("invalid or not implemented")
	}
}
