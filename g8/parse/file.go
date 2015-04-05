package parse

import (
	"io"

	"lonnie.io/e8vm/g8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseTopDecl(p *parser) ast.Decl {
	if p.SeeKeyword("const") {
		return parseConstDecls(p)
	} else if p.SeeKeyword("var") {
		return parseVarDecls(p)
	} else if p.SeeKeyword("func") {
		return parseFunc(p)
	} else if p.SeeKeyword("struct") {
		return parseStruct(p)
	}

	if len(p.Errs()) == 0 {
		// we only complain about this when there is no other error
		p.ErrorfHere("expect top level declaration")
	} else {
		p.Jail()
	}
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

// File parses a file.
func File(f string, r io.Reader) (*ast.File, []*lex8.Error) {
	p, _ := newParser(f, r)
	ret := parseFile(p)
	if es := p.Errs(); es != nil {
		return nil, es
	}
	return ret, nil
}
