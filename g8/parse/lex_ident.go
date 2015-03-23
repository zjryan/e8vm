package parse

import (
	"lonnie.io/e8vm/lex8"
)

func isLetter(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r == '_' {
		return true
	}
	return false
}

func lexIdent(x *lex8.Lexer) *lex8.Token {
	r := x.Rune()
	if !isLetter(r) {
		panic("must start with letter")
	}

	for {
		x.Next()
		r := x.Rune()
		if !isLetter(r) && !isDigit(r) {
			return x.MakeToken(Ident)
		}
	}
}
