package asm8

import (
	"lex8"
)

func lexString(x *lex8.Lexer) *lex8.Token {
	if !x.See('"') {
		panic("incorrect string start")
	}

	escaped := false

	for {
		x.Next()
		if x.Ended() {
			x.Err("unexpected eof in string")
			return x.MakeToken(String)
		} else if x.See('\n') {
			x.Err("unexpected endl in string")
			return x.MakeToken(String)
		}

		if escaped {
			escaped = false
		} else {
			if x.See('\\') {
				escaped = true
			} else if x.See('"') {
				x.Next()
				break
			}
		}
	}

	return x.MakeToken(String)
}
