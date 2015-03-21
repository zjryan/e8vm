package parse

import (
	"io"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseFile(p *parser) *ast.File {
	ret := new(ast.File)

	if p.SeeKeyword("import") {
		if imp := parseImports(p); imp != nil {
			ret.Imports = imp
		}
	}

	for !p.See(lex8.EOF) {
		if p.SeeKeyword("func") {
			if f := parseFunc(p); f != nil {
				ret.Decls = append(ret.Decls, f)
			}
		} else if p.SeeKeyword("var") {
			if v := parseVar(p); v != nil {
				ret.Decls = append(ret.Decls, v)
			}
		} else if p.SeeKeyword("const") {
			// TODO:
			p.ErrorfHere("const support not implemented yet")
			p.skipErrStmt()
		} else {
			p.ErrorfHere("expect top-declaration: func, var or const")
			return nil
		}

		p.BailOut()
	}

	return ret
}

// Result is a file parsing result
type Result struct {
	File   *ast.File
	Tokens []*lex8.Token
}

// FileResult returns a parsing result.
func FileResult(f string, rc io.ReadCloser) (*Result, []*lex8.Error) {
	p, rec := newParser(f, rc)
	parsed := parseFile(p)
	e := rc.Close()

	if e != nil {
		return nil, lex8.SingleErr(e)
	}
	if es := p.Errs(); es != nil {
		return nil, es
	}

	res := &Result{
		File:   parsed,
		Tokens: rec.Tokens(),
	}
	return res, nil
}

// File function parses a file into an AST.
func File(f string, rc io.ReadCloser) (*ast.File, []*lex8.Error) {
	res, es := FileResult(f, rc)
	if es != nil {
		return nil, es
	}
	return res.File, es
}
