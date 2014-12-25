package asm8

import (
	"lonnie.io/e8vm/lex8"
)

func lexComment(x *lex8.Lexer) *lex8.Token {
	if x.Rune() != '/' {
		panic("incorrect comment start")
	}

	x.Next()

	if x.Rune() == '/' {
		return lexLineComment(x)
	} else if x.Rune() == '*' {
		return lexBlockComment(x)
	}
	x.Err("illegal char %q", x.Rune())
	return x.MakeToken(lex8.Illegal)
}

func lexLineComment(x *lex8.Lexer) *lex8.Token {
	for {
		x.Next()
		if x.Ended() || x.Rune() == '\n' {
			break
		}
	}
	return x.MakeToken(Comment)
}

func lexBlockComment(x *lex8.Lexer) *lex8.Token {
	star := false
	for {
		x.Next()
		if x.Ended() {
			x.Err("unexpected eof in block comment")
			return x.MakeToken(Comment)
		}

		if star && x.Rune() == '/' {
			x.Next()
			break
		}

		star = x.Rune() == '*'
	}

	return x.MakeToken(Comment)
}
