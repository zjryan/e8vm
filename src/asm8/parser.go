package asm8

import (
	"io"
)

// Parser parses a file input stream into top-level syntax blocks.
type Parser struct {
	x    *Lexer
	errs *ErrList

	t            *Token
	inErr        bool
	parseComment bool

	parseFunc func(p *Parser) interface{}
}

// NewParser creates a new parser for parsring top-level syntax blocks.
func NewParser(file string, r io.ReadCloser) *Parser {
	ret := new(Parser)
	ret.x = NewLexer(file, r)
	ret.errs = NewErrList()

	return ret
}

func (p *Parser) err(pos *Pos, f string, args ...interface{}) {
	p.inErr = true
	p.errs.Addf(pos, f, args...)
}

func (p *Parser) skipLine(t *Token) {
	for t.Type != Endl && t.Type != EOF {
		t = p.x.Token()
	}
}

func (p *Parser) next() *Token {
	p.t = p.x.Token()
	if !p.parseComment {
		for p.t.Type == Comment {
			p.t = p.x.Token()
		}
	}
	return p.t
}

func typeStr(t int) string {
	switch t {
	case EOF:
		return "eof"
	case Comment:
		return "comment"
	case Keyword:
		return "keyword"
	case Operand:
		return "operand"
	case String:
		return "string"
	case Lbrace:
		return "'{'"
	case Rbrace:
		return "'}'"
	case Endl:
		return "end-line"
	case Illegal:
		return "illegal"
	}

	return "unknown"
}

func (p *Parser) expectKeyword(lit string) *Token {
	if p.inErr {
		return nil
	}

	if p.t.Type != Keyword || p.t.Lit != lit {
		p.errs.Addf(p.t.Pos, "expect keyword %s", lit)
		return nil
	}

	ret := p.t
	p.next()
	return ret
}

func (p *Parser) expect(t int) *Token {
	if p.inErr {
		return nil
	}

	if p.t.Type != t {
		p.errs.Addf(p.t.Pos, "expect %s, got %s",
			typeStr(t), typeStr(p.t.Type))
		return nil
	}

	ret := p.t
	p.next()
	return ret
}

func (p *Parser) acceptType(t int) bool {
	if p.t.Type != t {
		return false
	}
	p.next()
	return true
}

func (p *Parser) see(t int) bool {
	return p.t.Type == t
}

func (p *Parser) clearErr() {
	p.inErr = false
}

// Block returns the block by the parser function
func (p *Parser) Block() interface{} {
	if p.parseFunc == nil {
		return p.next()
	}

	return p.parseFunc(p)
}
