package parse

import (
	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseTopDecl(p *parser) ast.Decl {
	if p.SeeKeyword("const") {
		return parseConstDecls(p)
	} else if p.SeeKeyword("var") {
		return parseVarDecls(p)
	} else if p.SeeKeyword("func") {
		return parseFuncDecl(p)
	} else if p.SeeKeyword("struct") {
		return parseStructDecl(p)
	}

	p.ErrorfHere("expect top level declaration")
	p.Next() // make some progress anyway
	return nil
}

func parseFile(p *parser) *ast.File {
	var ret []ast.Decl
	for !p.See(lex8.EOF) {
		decl := parseTopDecl(p)
		if decl != nil {
			ret = append(ret, decl)
		}

		if p.InError() {
			p.skipErrStmt()
		}
	}

	return &ast.File{ret}
}
