package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func parseFile(p *parser) *file {
	ret := new(file)

	for p.seeKeyword("import") {
		panic("todo")
	}

	for !p.see(lex8.EOF) {
		if p.seeKeyword("func") {
			if f := parseFunc(p); f != nil {
				ret.Funcs = append(ret.Funcs, f)
			}
		} else if p.seeKeyword("var") {
			panic("todo")
		} else if p.seeKeyword("const") {
			panic("todo")
		} else {
			p.err(p.t.Pos, "expect top-declaration: func, var or const")
			return nil
		}
	}

	return ret
}
