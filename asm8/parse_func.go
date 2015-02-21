package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseFuncStmts(p *parser, f *ast.FuncDecl) {
	for !(p.see(Rbrace) || p.see(lex8.EOF)) {
		stmt := parseFuncStmt(p)
		if stmt != nil {
			f.Stmts = append(f.Stmts, stmt)
		}
		p.clearErr()
	}
}

func parseBareFunc(p *parser) *ast.FuncDecl {
	ret := new(ast.FuncDecl)
	ret.Name = &lex8.Token{
		Type: Operand,
		Lit:  "_",
		Pos:  nil,
	}
	parseFuncStmts(p, ret)
	return ret
}

func parseFunc(p *parser) *ast.FuncDecl {
	ret := new(ast.FuncDecl)

	ret.Kw = p.expectKeyword("func")
	ret.Name = p.expect(Operand)

	if ret.Name != nil {
		name := ret.Name.Lit
		if !isIdent(name) {
			p.err(ret.Name.Pos, "invalid func name %q", name)
		}
	}

	ret.Lbrace = p.expect(Lbrace)
	if p.skipErrStmt() {
		return ret
	}

	parseFuncStmts(p, ret)

	ret.Rbrace = p.expect(Rbrace)
	ret.Semi = p.expect(Semi)
	p.skipErrStmt()

	return ret
}
