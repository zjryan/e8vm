package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Func is an assembly function.
type Func struct {
	stmts []*stmt

	kw, name             *lex8.Token
	lbrace, rbrace, semi *lex8.Token

	addr uint32
}

func (f *Func) parseStmts(p *Parser) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseStmt(p)
		if stmt != nil {
			f.stmts = append(f.stmts, stmt)
		}
	}
}

func parseBareFunc(p *Parser) *Func {
	ret := new(Func)
	ret.name = &lex8.Token{Operand, "_", nil}
	ret.parseStmts(p)
	return ret
}

func parseFunc(p *Parser) *Func {
	ret := new(Func)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)

	name := ret.name.Lit
	if !isIdent(name) {
		p.err(ret.name.Pos, "invalid function name %q", name)
	}

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
