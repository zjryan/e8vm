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

	indices map[string]uint32
	pkgUsed map[string]struct{}
}

func newBuilder() *builder {
	ret := new(builder)
	ret.errs = lex8.NewErrList()
	ret.scope = newSymScope()
	ret.indices = make(map[string]uint32)
	ret.pkgUsed = make(map[string]struct{})

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

func (b *builder) index(name string, index uint32) {
	_, found := b.indices[name]
	if found {
		panic("redeclare")
	}

	b.indices[name] = index
}

func (b *builder) getIndex(name string) uint32 {
	ret, found := b.indices[name]
	if !found {
		panic("not found")
	}
	return ret
}
