package g8

import (
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/g8/types"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

type builder struct {
	*lex8.ErrorList
	path string

	p         *ir.Pkg
	f         *ir.Func
	fretNamed bool
	fretRef   *ref

	b     *ir.Block
	scope *sym8.Scope

	builtin uint32 // the index of the builtin package

	exprFunc func(b *builder, expr ast.Expr) *ref
	stmtFunc func(b *builder, stmt ast.Stmt)
	irLog    io.WriteCloser
}

func newBuilder(path string) *builder {
	ret := new(builder)
	ret.ErrorList = lex8.NewErrorList()
	ret.path = path
	ret.p = ir.NewPkg(path)
	ret.scope = sym8.NewScope() // package scope

	return ret
}

func (b *builder) newTemp(t types.T) *ref {
	return newRef(t, b.f.NewTemp(t.Size()))
}

func (b *builder) newLocal(t types.T, name string) ir.Ref {
	return b.f.NewLocal(t.Size(), name)
}

func (b *builder) buildExpr(expr ast.Expr) *ref {
	if b.exprFunc != nil {
		return b.exprFunc(b, expr)
	}
	return nil
}

func (b *builder) buildStmts(stmts []ast.Stmt) {
	if b.stmtFunc == nil {
		return
	}

	for _, stmt := range stmts {
		b.stmtFunc(b, stmt)
	}
}
