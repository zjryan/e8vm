package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

// Parser parses a file input stream into top-level syntax blocks.
type parser struct {
	x *stmtLexer
	*lex8.Parser
}

var types = func() *lex8.Types {
	ret := lex8.NewTypes()
	o := func(t int, name string) {
		ret.Register(t, name)
	}

	o(Keyword, "keyword")
	o(Operand, "operand")
	o(Lbrace, "'{'")
	o(Rbrace, "'}'")
	o(Semi, "';'")
	o(Endl, "end-line")

	return ret
}()

func newParser(file string, r io.Reader) *parser {
	ret := new(parser)
	ret.x = newStmtLexer(file, r)
	ret.Parser = lex8.NewParser(ret.x, types)
	return ret
}

func (p *parser) SeeKeyword(kw string) bool {
	return p.SeeLit(Keyword, kw)
}

func (p *parser) ExpectKeyword(kw string) *lex8.Token {
	return p.ExpectLit(Keyword, kw)
}

func (p *parser) skipErrStmt() bool {
	return p.SkipErrStmt(Semi)
}

func (p *parser) comments() []*lex8.Token {
	return p.x.Comments()
}
