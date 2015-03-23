package parse

import (
	"lonnie.io/e8vm/lex8"
)

func lexOperator(x *lex8.Lexer, r rune) *lex8.Token {
	switch r {
	case '{', '}', '(', ')', '[', ']', ',', ';':
		return x.MakeToken(Operator)
	case '/':
		r2 := x.Rune()
		if r2 == '/' || r2 == '*' {
			return lex8.LexComment(x)
		} else if r2 == '=' {
			x.Next()
			return x.MakeToken(Operator)
		} else {
			return x.MakeToken(Operator)
		}
	case '+', '-', '&', '|':
		r2 := x.Rune()
		if r2 == r || r2 == '=' {
			x.Next()
		}
		return x.MakeToken(Operator)
	case '*', '%', '^', '=', '!', ':':
		r2 := x.Rune()
		if r2 == '=' {
			x.Next()
		}
		return x.MakeToken(Operator)
	case '.':
		r2 := x.Rune()
		if r2 == '.' {
			x.Next()
			r3 := x.Rune()
			if r3 != '.' {
				x.Errorf("expect ..., but see ..")
				return x.MakeToken(Operator)
			}
			x.Next()
		}
		return x.MakeToken(Operator)
	case '>', '<':
		r2 := x.Rune()
		if r2 == r {
			x.Next()
			r3 := x.Rune()
			if r3 == '=' {
				x.Next()
			}
		} else if r2 == '=' {
			x.Next()
		}
		return x.MakeToken(Operator)
	}

	return nil
}
