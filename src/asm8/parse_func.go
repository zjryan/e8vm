package asm8

import (
	"lex8"
)

// Func is an assembly function.
type Func struct {
	stmts []*stmt

	kw, name, lbrace, rbrace, semi *lex8.Token
}

func (f *Func) parseStmts(p *Parser) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseStmt(p)
		if stmt != nil {
			f.stmts = append(f.stmts, stmt)
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

	ret.parseStmts(p)

	ret.rbrace = p.expect(Rbrace)
	ret.semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
