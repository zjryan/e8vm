package parse

import (
	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseImportStmt(p *parser) *ast.ImportStmt {
	ret := new(ast.ImportStmt)

	path := p.Expect(String)
	ret.Path = path

	if path != nil {
		if p.See(Operand) {
			ret.As = p.Shift()
		}
	}

	p.Expect(Semi)
	if p.skipErrStmt() {
		return nil
	}

	return ret
}

func parseImports(p *parser) *ast.ImportDecl {
	ret := new(ast.ImportDecl)
	ret.Kw = p.ExpectKeyword("import")
	ret.Lbrace = p.Expect(Lbrace)
	if p.skipErrStmt() { // header broken
		return ret
	}

	for !p.See(Rbrace) && !p.See(lex8.EOF) {
		imp := parseImportStmt(p)
		if imp != nil {
			ret.Stmts = append(ret.Stmts, imp)
		}
	}

	ret.Rbrace = p.Expect(Rbrace)
	ret.Semi = p.Expect(Semi)
	p.skipErrStmt()

	return ret
}
