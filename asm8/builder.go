package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Builder manipulates an AST, checks its syntax, and builds the assembly
type Builder struct {
	errs  *lex8.ErrorList
	scope *SymScope

	curPkg *pkgObj

	hasError bool
}

func newBuilder() *Builder {
	ret := new(Builder)
	ret.errs = lex8.NewErrList()
	ret.scope = NewSymScope()

	return ret
}

func (b *Builder) err(p *lex8.Pos, f string, args ...interface{}) {
	b.hasError = true
	b.errs.Addf(p, f, args...)
}

// Errs returns the building errors.
func (b *Builder) Errs() []*lex8.Error {
	return b.errs.Errs
}

func (b *Builder) clearErr() {
	b.hasError = false
}
