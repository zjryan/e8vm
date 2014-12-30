package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func parseOps(p *parser) (ops []*lex8.Token) {
	for !p.acceptType(Semi) {
		t := p.expect(Operand)
		if t != nil {
			ops = append(ops, t)
		} else {
			ops = nil // error now
			if p.see(lex8.EOF) {
				break
			}
			p.next() // proceed anyway for other stuff
		}
	}

	p.clearErr()
	return ops
}
