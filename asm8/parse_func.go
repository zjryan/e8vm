package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func parseFuncStmts(p *parser, f *funcDecl) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseFuncStmt(p)
		if stmt != nil {
			f.stmts = append(f.stmts, stmt)
		}
		p.clearErr()
	}
}

func parseBareFunc(p *parser) *funcDecl {
	ret := new(funcDecl)
	ret.name = &lex8.Token{Operand, "_", nil}
	parseFuncStmts(p, ret)
	return ret
}

func parseFunc(p *parser) *funcDecl {
	ret := new(funcDecl)

	ret.kw = p.expectKeyword("func")
	ret.name = p.expect(Operand)

	name := ret.name.Lit
	if !isIdent(name) {
		p.err(ret.name.Pos, "invalid func name %q", name)
	}

	ret.lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	parseFuncStmts(p, ret)

	ret.rbrace = p.expect(Rbrace)
	ret.semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
