package asm8

import (
	"io"

	"lex8"
)

// Parser parses a file input stream into top-level syntax blocks.
type Parser struct {
	x    *StmtLexer
	errs *lex8.ErrorList

	t     *lex8.Token
	inErr bool
}

func newParser(file string, r io.ReadCloser) *Parser {
	ret := new(Parser)
	ret.x = NewStmtLexer(file, r)
	ret.errs = lex8.NewErrList()
	ret.next()

	return ret
}

func (p *Parser) err(pos *lex8.Pos, f string, args ...interface{}) {
	p.inErr = true
	p.errs.Addf(pos, f, args...)
}

func (p *Parser) see(t int) bool { return p.t.Type == t }
func (p *Parser) seeKeyword(kw string) bool {
	return p.see(Keyword) && p.t.Lit == kw
}
func (p *Parser) clearErr()          { p.inErr = false }
func (p *Parser) hasErr() bool       { return p.inErr }
func (p *Parser) token() *lex8.Token { return p.t }

func (p *Parser) next() *lex8.Token {
	p.t = p.x.Token()
	return p.t
}

func typeStr(t int) string {
	switch t {
	case lex8.EOF:
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
	case Semi:
		return ";"
	case Endl:
		return "end-line"
	case lex8.Illegal:
		return "illegal"
	}

	return "unknown"
}

func (p *Parser) expectKeyword(lit string) *lex8.Token {
	if p.inErr {
		return nil
	}

	if p.t.Type != Keyword || p.t.Lit != lit {
		p.err(p.t.Pos, "expect keyword %s, got %s", lit, typeStr(p.t.Type))
		return nil
	}

	ret := p.t
	p.next()
	return ret
}

func (p *Parser) expect(t int) *lex8.Token {
	if p.inErr {
		return nil
	}

	if p.t.Type != t {
		p.err(p.t.Pos, "expect %s, got %s", typeStr(t), typeStr(p.t.Type))
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

func (p *Parser) skipStmt() {
	for !(p.see(Semi) || p.see(lex8.EOF)) {
		p.next()
	}

	if p.see(Semi) {
		p.next()
	}
}

func (p *Parser) skipErrStmt() bool {
	if !p.inErr {
		return false
	}

	p.skipStmt()
	p.clearErr()
	return true
}

// Errs returns the parsing errors
func (p *Parser) Errs() []*lex8.Error {
	ret := p.x.Errs()
	if ret != nil {
		return ret
	}
	return p.errs.Errs
}
