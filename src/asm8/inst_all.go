package asm8

import (
	"lex8"
)

func makeInst(p *Parser, ops []*lex8.Token) (i *inst) {
	var hit bool

	return nil

	if i, hit = makeInstReg(p, ops); hit {
		return i
	}
	if i, hit = makeInstImm(p, ops); hit {
		return i
	}
	if i, hit = makeInstBr(p, ops); hit {
		return i
	}
	if i, hit = makeInstSys(p, ops); hit {
		return i
	}

	return nil
}
