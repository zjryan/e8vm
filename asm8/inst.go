package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type inst struct {
	inst   uint32
	pkg    string
	sym    string
	fill   int
	symTok *lex8.Token
}

type instResolver func(lex8.Logger, []*lex8.Token) (*inst, bool)
type instResolvers []instResolver

func (rs instResolvers) resolve(log lex8.Logger, ops []*lex8.Token) *inst {
	for _, r := range rs {
		if i, hit := r(log, ops); hit {
			return i
		}
	}

	op0 := ops[0]
	log.Errorf(op0.Pos, "invalid asm instruction %q", op0.Lit)
	return nil
}
