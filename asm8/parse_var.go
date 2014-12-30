package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func parseVarStmts(p *parser, v *varDecl) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseVarStmt(p)
		if stmt != nil {
			v.stmts = append(v.stmts, stmt)
		}
		p.clearErr()
	}
}

func parseVar(p *parser) *varDecl {
	ret := new(varDecl)

	ret.kw = p.expectKeyword("var")
	ret.name = p.expect(Operand)

	if ret.name != nil {
		name := ret.name.Lit
		if !isIdent(name) {
			p.err(ret.name.Pos, "invalid var name %q", name)
		}
	}

	ret.lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	parseVarStmts(p, ret)

	ret.rbrace = p.expect(Rbrace)
	ret.semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
