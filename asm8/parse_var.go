package asm8

func parseVarStmts(p *parser, v *varDecl) {

}

func parseVar(p *parser) *varDecl {
	ret := new(varDecl)

	ret.kw = p.expectKeyword("var")
	ret.name = p.expect(Operand)

	name := ret.name.Lit
	if !isIdent(name) {
		p.err(ret.name.Pos, "invalid var name %q", name)
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
