package asm8

import (
	"lex8"
)

func parseInst(p *Parser, ops []*lex8.Token) (i *inst) {
	var hit bool

	if i, hit = parseInstReg(p, ops); hit {
		return i
	}
	if i, hit = parseInstImm(p, ops); hit {
		return i
	}
	if i, hit = parseInstBr(p, ops); hit {
		return i
	}
	if i, hit = parseInstSys(p, ops); hit {
		return i
	}

	op0 := ops[0]
	p.err(op0.Pos, "invalid asm instruction %q", op0.Lit)
	return nil
}
