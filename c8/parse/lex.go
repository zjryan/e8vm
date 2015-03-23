package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func lexC8(x *lex8.Lexer) *lex8.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '\n':
		x.Next()
		return x.MakeToken(Endl)
	case '{', '}', '(', ')', '[', ']':
		x.Next()
		return x.MakeToken(Operator)
	case '/':
		x.Next()
		r2 := x.Rune()
		if r2 == '/' || r2 == '*' {
			return lex8.LexComment(x)
		} else if r2 == '=' {
			x.Next()
			return x.MakeToken(Operator)
		} else {
			return x.MakeToken(Operator)
		}
	case '"':
		return lex8.LexString(x, String)
	}

	if isDigit(r) {
		panic("todo: lex number")
	}

	x.Errorf("illegal char %q", r)
	x.Next()
	return x.MakeToken(lex8.Illegal)
}

// NewLexer creates a new c8 lexer for a file input stream.
func newLexer(file string, r io.Reader) *lex8.Lexer {
	ret := lex8.NewLexer(file, r)
	ret.LexFunc = lexC8
	return ret
}

// Tokens parses a file into a token array
func Tokens(f string, r io.Reader) ([]*lex8.Token, []*lex8.Error) {
	x := newLexer(f, r)
	toks := lex8.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, errs
	}
	return toks, nil
}
