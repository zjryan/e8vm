package asm8

import (
	"lex8"
)

// Func is an assembly function.
type Func struct {
	Insts []*Inst

	kw, name, lbrace, rbrace, semi *lex8.Token
}

func (f *Func) parseInsts(p *Parser) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		inst := parseInst(p)
		if inst != nil {
			f.Insts = append(f.Insts, inst)
		}
	}
}

func parseFunc(p *Parser) *Func {
	ret := new(Func)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)
	ret.lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	ret.parseInsts(p)

	ret.rbrace = p.expect(Rbrace)
	ret.semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
