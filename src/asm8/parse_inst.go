package asm8

import (
	"lex8"
)

// Inst is an assembly instruction line.
type Inst struct {
	Ops []*lex8.Token

	label  string
	inst   uint32
	symbol string
	fill   int
}

func parseInst(p *Parser) *Inst {
	ret := new(Inst)

	// a good assembly instruction is a series of ops that ends with
	// a semicolon or a right-brace
	for {
		if p.acceptType(Semi) || p.see(Rbrace) || p.see(lex8.EOF) {
			break
		}

		if p.see(Lbrace) {
			p.expect(Operand)
			return nil
		}

		t := p.expect(Operand)
		if t != nil {
			ret.Ops = append(ret.Ops, t)
		} else {
			ret = nil // error now
			if p.see(Lbrace) {
				break
			}
			p.next() // proceed anyway for other stuff
		}
	}

	p.clearErr()

	if ret == nil {
		return nil
	}

	return ret
}
