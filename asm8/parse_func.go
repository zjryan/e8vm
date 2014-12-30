package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func (f *funcDecl) parseStmts(p *Parser) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseStmt(p)
		if stmt != nil {
			f.stmts = append(f.stmts, stmt)
		}
	}
}

func parseBareFunc(p *Parser) *funcDecl {
	ret := new(funcDecl)
	ret.name = &lex8.Token{Operand, "_", nil}
	ret.parseStmts(p)
	return ret
}

func parseFunc(p *Parser) *funcDecl {
	ret := new(funcDecl)

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
