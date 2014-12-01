package asm8

import (
	"io"
)

// Parser parses a file input stream into top-level syntax blocks.
type Parser struct {
	x    *Lexer
	errs *ErrList
}

// NewParser creates a new parser for parsring top-level syntax blocks.
func NewParser(file string, r io.ReadCloser) *Parser {
	ret := new(Parser)
	ret.x = NewLexer(file, r)
	ret.errs = NewErrList()

	return ret
}

func (p *Parser) err(pos *Pos, f string, args ...interface{}) {
	p.errs.Addf(pos, f, args...)
}

func (p *Parser) parseFunc() interface{} {
	panic("todo")
}

func (p *Parser) parseVar() interface{} {
	panic("todo")
}

func (p *Parser) parseImport() interface{} {
	panic("todo")
}

func (p *Parser) parseConst() interface{} {
	panic("todo")
}

func (p *Parser) skipLine(t *Token) {
	for t.Type != Endl && t.Type != EOF {
		t = p.x.Token()
	}
}

func (p *Parser) Block() interface{} {
	t := p.x.Token()
	if t.Type != Keyword {
		p.err(t.Pos, "expect top-level declaration")
		p.skipLine(t)
		return nil
	}

	switch t.Lit {
	case "func":
		return p.parseFunc()
	case "var":
		return p.parseVar()
	case "const":
		return p.parseConst()
	case "import":
		return p.parseImport()
	default:
		p.err(t.Pos, "unknown top-level declaration keyword")
		p.skipLine(t)
		return nil
	}

	return nil
}
