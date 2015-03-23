package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

func lexC8(x *lex8.Lexer) *lex8.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '\n':
		x.Next()
		return x.MakeToken(Endl)
	case '"':
		return lex8.LexString(x, String)
	case '\'':
		panic("char is on todo")
	}

	if isDigit(r) {
		return lexNumber(x)
	} else if isLetter(r) {
		return lexIdent(x)
	}

	// always make progress at this point
	x.Next()
	t := lexOperator(x, r)
	if t != nil {
		return t
	}

	x.Errorf("illegal char %q", r)
	return x.MakeToken(lex8.Illegal)
}

// NewLexer creates a new c8 lexer for a file input stream.
func newLexer(file string, r io.Reader) *lex8.Lexer {
	ret := lex8.NewLexer(file, r)
	ret.LexFunc = lexC8
	return ret
}
