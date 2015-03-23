package parse

import (
	"fmt"
	"io"

	"lonnie.io/e8vm/lex8"
)

type parser struct {
	x lex8.Tokener
	*lex8.Parser
}

func makeTokener(f string, r io.Reader) (lex8.Tokener, *lex8.Recorder) {
	var x lex8.Tokener = newLexer(f, r)
	x = newSemiInserter(x)
	rec := lex8.NewRecorder(x)
	return lex8.NewCommentRemover(rec), rec
}

func newParser(f string, r io.Reader) (*parser, *lex8.Recorder) {
	ret := new(parser)
	x, rec := makeTokener(f, r)
	ret.x = x
	ret.Parser = lex8.NewParser(ret.x, Types)
	return ret, rec
}

func (p *parser) SeeOp(ops ...string) bool {
	t := p.Token()
	if t.Type != Operator {
		return false
	}
	for _, op := range ops {
		if t.Lit == op {
			return true
		}
	}
	return false
}

func (p *parser) typeStr(t *lex8.Token) string {
	if t.Type == Operator {
		return fmt.Sprintf("'%s'", t.Lit)
	}
	return TypeStr(t.Type)
}

func (p *parser) ExpectOp(op string) *lex8.Token {
	t := p.Token()
	if t.Type != Operator || t.Lit != op {
		p.ErrorfHere("expect '%s', got %s", t.Lit, p.typeStr(t))
		return nil
	}

	return p.Shift()
}

// Tokens parses a file into a token array
func Tokens(f string, r io.Reader) ([]*lex8.Token, []*lex8.Error) {
	x, _ := makeTokener(f, r)
	toks := lex8.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, errs
	}
	return toks, nil
}
