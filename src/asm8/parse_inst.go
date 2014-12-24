package asm8

import (
	"lex8"
)

type instParse func(*Parser, []*lex8.Token) (*inst, bool)
type instParsers []instParse

func (ips instParsers) parse(p *Parser, ops []*lex8.Token) *inst {
	for _, ip := range ips {
		if i, hit := ip(p, ops); hit {
			return i
		}
	}

	op0 := ops[0]
	p.err(op0.Pos, "invalid asm instruction %q", op0.Lit)
	return nil
}
