package g8

import (
	"bytes"
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

func buildOperand(b *builder, op *ast.Operand) ir.Ref {
	switch op.Token.Type {
	case parse.Int:
		panic("todo: integer")
	default:
		panic("invalid or not implemented")
	}
}

func buildExpr(b *builder, expr ast.Expr) ir.Ref {
	if expr == nil {
		return nil
	}

	switch expr := expr.(type) {
	default:
		_ = expr
		panic("invalid or not implemented")
	}
}

func buildExprList(b *builder, list *ast.ExprList) []ir.Ref {
	ret := make([]ir.Ref, 0, list.Len())
	for _, expr := range list.Exprs {
		ref := buildExpr(b, expr)
		if ref == nil {
			return nil
		}
		ret = append(ret, ref)
	}
	return ret
}

func buildStmt(b *builder, stmt ast.Stmt) {
	switch stmt := stmt.(type) {
	case *ast.DefineStmt:
		if stmt.Left.Len() == stmt.Right.Len() {
			rights := buildExprList(b, stmt.Right)
			if rights == nil {
				return
			}

			// TODO: decalre left as new variables based on type on rights
			lefts := buildExprList(b, stmt.Left)
			if lefts == nil {
				return
			}

			// TODO: check type matching

			for i, left := range lefts {
				b.b.Assign(left, rights[i])
			}
		} else if stmt.Right.Len() == 1 {
			panic("todo: right might be a function call that retunrs a list")
		} else {
			b.Errorf(stmt.Define.Pos, "mismatch definition")
		}
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
				b.b.Assign(left, rights[i])
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

func buildBareFunc(b *builder, stmts []ast.Stmt) *link8.Pkg {
	b.f = b.p.NewFunc("main")
	b.b = b.f.NewBlock()

	for _, stmt := range stmts {
		buildStmt(b, stmt)
	}

	return ir.BuildPkg(b.p)
}

// BuildBareFunc builds a bare main function of signature func main()
func BuildBareFunc(f string, r io.Reader) ([]byte, []*lex8.Error) {
	stmts, es := parse.Stmts(f, r)
	if es != nil {
		return nil, es
	}

	b := newBuilder("_")
	pkg := buildBareFunc(b, stmts)
	if es := b.Errs(); es != nil {
		return nil, es
	}

	buf := new(bytes.Buffer)
	e := link8.LinkMain(pkg, buf)
	if e != nil {
		return nil, lex8.SingleErr(e)
	}

	return buf.Bytes(), nil
}
