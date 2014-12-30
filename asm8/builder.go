package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Builder manipulates an AST, checks its syntax, and builds the assembly
type builder struct {
	errs  *lex8.ErrorList
	scope *symScope

	curPkg *lib

	hasError bool
}

func newBuilder() *builder {
	ret := new(builder)
	ret.errs = lex8.NewErrList()
	ret.scope = newSymScope()

	return ret
}

func (b *builder) err(p *lex8.Pos, f string, args ...interface{}) {
	b.hasError = true
	b.errs.Addf(p, f, args...)
}

// Errs returns the building errors.
func (b *builder) Errs() []*lex8.Error {
	return b.errs.Errs
}

func (b *builder) clearErr() {
	b.hasError = false
}
