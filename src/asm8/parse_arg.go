package asm8

import (
	"lex8"
)

func argCount(p *Parser, ops []*lex8.Token, n int) bool {
	if len(ops) == n+1 {
		return true
	}

	p.err(ops[0].Pos, "%q needs %d arguments", ops[0].Lit, n)
	return false
}
