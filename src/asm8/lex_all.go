package asm8

import (
	"io"
)

func lexAsm8(x *Lexer) *Token {
	if x.isWhite(x.r) {
		panic("incorrect token start")
	}

	switch x.r {
	case '\n':
		x.next()
		return x.token(Endl)
	case '{':
		x.next()
		return x.token(Lbrace)
	case '}':
		x.next()
		return x.token(Rbrace)
	case '/':
		return lexComment(x)
	case '"':
		return lexString(x)
	}

	if isOperandChar(x.r) {
		return lexOperand(x)
	}

	x.err("illegal char %q", x.r)
	x.next()
	return x.token(Illegal)
}

// NewLexer creates a new lexer of a file stream.
func NewLexer(file string, r io.ReadCloser) *Lexer {
	ret := newLexer(file, r)
	ret.lexFunc = lexAsm8
	return ret
}
