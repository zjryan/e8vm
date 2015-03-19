package parse

import (
	"lonnie.io/e8vm/lex8"
)

func isOperandChar(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}

	for _, x := range []rune{'_', '+', '-', '.', ':'} {
		if r == x {
			return true
		}
	}

	return false
}

func isKeyword(lit string) bool {
	switch lit {
	case "func", "var", "const", "import":
		return true
	}
	return false
}

func lexOperand(x *lex8.Lexer) *lex8.Token {
	if !isOperandChar(x.Rune()) {
		panic("incorrect operand start")
	}

	for {
		x.Next()
		if x.Ended() || !isOperandChar(x.Rune()) {
			break
		}
	}

	ret := x.MakeToken(Operand)
	if isKeyword(ret.Lit) {
		ret.Type = Keyword
	}
	return ret
}
