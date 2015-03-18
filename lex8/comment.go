package lex8

// LexComment lexes a c style comment
func LexComment(x *Lexer) *Token {
	if x.Rune() != '/' {
		panic("incorrect comment start")
	}

	x.Next()

	if x.Rune() == '/' {
		return lexLineComment(x)
	} else if x.Rune() == '*' {
		return lexBlockComment(x)
	}
	x.Errorf("illegal char %q", x.Rune())
	return x.MakeToken(Illegal)
}

func lexLineComment(x *Lexer) *Token {
	for {
		x.Next()
		if x.Ended() || x.Rune() == '\n' {
			break
		}
	}
	return x.MakeToken(Comment)
}

func lexBlockComment(x *Lexer) *Token {
	star := false
	for {
		x.Next()
		if x.Ended() {
			x.Errorf("unexpected eof in block comment")
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
