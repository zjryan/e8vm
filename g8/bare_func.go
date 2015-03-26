package g8

import (
	"bytes"
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

type builder struct {
	*lex8.ErrorList
}

func newBuilder() *builder {
	ret := new(builder)
	ret.ErrorList = lex8.NewErrorList()
	return ret
}

func buildBareFunc(b *builder, stmts []ast.Stmt) *link8.Pkg {
	panic("todo")
}

func BuildBareFunc(f string, r io.Reader) ([]byte, []*lex8.Error) {
	stmts, es := parse.Stmts(f, r)
	if es != nil {
		return nil, es
	}

	b := newBuilder()
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
