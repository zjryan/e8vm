package asm8

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
		return lexComment(x)
	case '"':
		return lexString(x)
	}

	if isOperandChar(r) {
		return lexOperand(x)
	}

	x.Err("illegal char %q", r)
	x.Next()
	return x.MakeToken(lex8.Illegal)
}

// NewLexer creates a new lexer of a file stream.
func NewLexer(file string, r io.ReadCloser) *lex8.Lexer {
	ret := lex8.NewLexer(file, r)
	ret.LexFunc = lexAsm8
	return ret
}

// Tokens parses a file in a token array
func Tokens(f string, r io.ReadCloser) ([]*lex8.Token, []*lex8.Error) {
	x := NewLexer(f, r)

	var ret []*lex8.Token

	for {
		t := x.Token()
		ret = append(ret, t)
		if t.Type == lex8.EOF {
			break
		}
	}

	errs := x.Errs()
	if errs != nil {
		return nil, errs
	}

	return ret, nil
}
