package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type resolver struct {
	errs *lex8.ErrorList
}

func newResolver() *resolver {
	ret := new(resolver)
	ret.errs = lex8.NewErrorList()

	return ret
}

func (r *resolver) Errorf(pos *lex8.Pos, fmt string, args ...interface{}) {
	// simply relay, make it a logger
	r.errs.Errorf(pos, fmt, args...)
}

func (r *resolver) Errs() []*lex8.Error {
	return r.errs.Errs()
}
