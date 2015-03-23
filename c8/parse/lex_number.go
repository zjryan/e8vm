package parse

import (
	"lonnie.io/e8vm/lex8"
)

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func lexNumber(x *lex8.Lexer) *lex8.Token {
	// TODO: lex floating point as well

	r := x.Rune()
	if !isDigit(r) {
		panic("not starting with a number")
	}

	for {
		x.Next()
		r := x.Rune()
		if !isDigit(r) {
			return x.MakeToken(Int)
		}
	}
}
