package parse

import (
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

// Tokens parses a file into a token array
func Tokens(f string, r io.Reader) ([]*lex8.Token, []*lex8.Error) {
	x, _ := makeTokener(f, r)
	toks := lex8.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, errs
	}
	return toks, nil
}
