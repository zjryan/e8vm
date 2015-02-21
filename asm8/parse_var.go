package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseVarStmts(p *parser, v *ast.VarDecl) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseVarStmt(p)
		if stmt != nil {
			v.Stmts = append(v.Stmts, stmt)
		}
		p.clearErr()
	}
}

func parseVar(p *parser) *ast.VarDecl {
	ret := new(ast.VarDecl)

	ret.Kw = p.expectKeyword("var")
	ret.Name = p.expect(Operand)

	if ret.Name != nil {
		name := ret.Name.Lit
		if !isIdent(name) {
			p.err(ret.Name.Pos, "invalid var name %q", name)
		}
	}

	ret.Lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	parseVarStmts(p, ret)

	ret.Rbrace = p.expect(Rbrace)
	ret.Semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
