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

func isKeyword(lit string) bool {
	switch lit {
	case "func", "var", "const", "import", "type",
		"if", "else", "for",
		"break", "continue", "return",
		"switch", "case", "default", "fallthrough",
		"range", "struct", "interface", "goto":
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
			break
		}
	}

	ret := x.MakeToken(Ident)
	if isKeyword(ret.Lit) {
		ret.Type = Keyword
	}
	return ret
}
