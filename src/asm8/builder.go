package asm8

import (
	"io"

	"lex8"
)

// Builder manipulates an AST, checks its syntax, and builds the assembly
type Builder struct {
	p     *Parser
	errs  *lex8.ErrorList
	scope *SymScope
}

func newBuilder(file string, r io.ReadCloser) *Builder {
	ret := new(Builder)
	ret.p = newParser(file, r)
	ret.errs = lex8.NewErrList()
	ret.scope = NewSymScope()

	return ret
}

// Err reports a building/semantics error.
func (b *Builder) err(p *lex8.Pos, f string, args ...interface{}) {
	b.errs.Addf(p, f, args...)
}
