package parse

import (
	"io"

	"lonnie.io/e8vm/lex8"
)

func lexAsm8(x *lex8.Lexer) *lex8.Token {
	r := x.Rune()
	if x.IsWhite(r) {
		panic("incorrect token start")
	}

	switch r {
	case '\n':
		x.Next()
		return x.MakeToken(Endl)
	case '{':
		x.Next()
		return x.MakeToken(Lbrace)
	case '}':
		x.Next()
		return x.MakeToken(Rbrace)
	case '/':
		x.Next()
		return lex8.LexComment(x)
	case '"':
		return lex8.LexString(x, String, '"')
	}

	if isOperandChar(r) {
		return lexOperand(x)
	}

	x.Errorf("illegal char %q", r)
	x.Next()
	return x.MakeToken(lex8.Illegal)
}

// NewLexer creates a new lexer of a file stream.
func newLexer(file string, r io.Reader) *lex8.Lexer {
	ret := lex8.NewLexer(file, r)
	ret.LexFunc = lexAsm8
	return ret
}

// Tokens parses a file in a token array
func Tokens(f string, r io.Reader) ([]*lex8.Token, []*lex8.Error) {
	x := newLexer(f, r)
	toks := lex8.TokenAll(x)
	if es := x.Errs(); es != nil {
		return nil, es
	}
	return toks, nil
}
