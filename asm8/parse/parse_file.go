package parse

import (
	"io"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func parseFile(p *parser) *ast.File {
	ret := new(ast.File)

	for !p.see(lex8.EOF) {
		if p.seeKeyword("func") {
			if f := parseFunc(p); f != nil {
				ret.Funcs = append(ret.Funcs, f)
			}
		} else if p.seeKeyword("var") {
			if v := parseVar(p); v != nil {
				ret.Vars = append(ret.Vars, v)
			}
		} else if p.seeKeyword("const") {
			// TODO:
			p.err(p.t.Pos, "const support not implemented yet")
			p.skipErrStmt()
		} else {
			p.err(p.t.Pos, "expect top-declaration: func, var or const")
			return nil
		}

		p.clearErr()
	}

	return ret
}

// Parse parses a file into an AST.
func File(f string, rc io.ReadCloser) (*ast.File, []*lex8.Error) {
	p := newParser(f, rc)
	parsed := parseFile(p)
	e := rc.Close()

	if e != nil {
		return nil, lex8.SingleErr(e)
	}
	if es := p.Errs(); es != nil {
		return nil, es
	}
	return parsed, nil
}
