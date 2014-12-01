package asm8

import (
	"io"
)

// NewParser creates a new parser for parsring top-level syntax blocks.
func NewParser(file string, r io.ReadCloser) *Parser {
	ret := NewParser(file, r)
	ret.ParseFunc = parseAsm8
	return ret
}

func parseAsm8(p *Parser) interface{} {
	if p.seeKeyword("func") {
		return parseFunc(p)
	}

	p.err(p.t.Pos, "expect top-declaration")
	return int(0) // place holder for error
}

// Blocks returns the AST of a file input stream or parsing errors.
func Blocks(f string, r io.ReadCloser) ([]interface{}, []*Error) {
	var ret []interface{}
	p := NewParser(f, r)

	for {
		b := p.Block()
		if b == nil {
			break
		}
		ret = append(ret, b)
	}

	errs := p.Errs()
	if errs != nil {
		return nil, errs
	}
	return ret, nil
}
