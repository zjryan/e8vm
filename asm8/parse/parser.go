package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

// Parser parses a file input stream into top-level syntax blocks.
type parser struct {
	x lex8.Tokener
	*lex8.Parser
}

func newParser(f string, r io.Reader) (*parser, *lex8.Recorder) {
	ret := new(parser)

	var x lex8.Tokener = newLexer(f, r)
	x = newSemiInserter(x)
	rec := lex8.NewRecorder(x)
	ret.x = lex8.NewCommentRemover(rec)
	ret.Parser = lex8.NewParser(ret.x, Types)
	return ret, rec
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
