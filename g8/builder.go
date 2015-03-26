package g8

import (
	"lonnie.io/e8vm/g8/ir"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

type builder struct {
	*lex8.ErrorList
	path string

	p     *ir.Pkg
	f     *ir.Func
	b     *ir.Block
	scope *sym8.Scope
}

func newBuilder(path string) *builder {
	ret := new(builder)
	ret.ErrorList = lex8.NewErrorList()
	ret.path = path
	ret.p = ir.NewPkg(path)
	ret.scope = sym8.NewScope() // package scope

	return ret
}
