package asm8

func lexComment(x *Lexer) *Token {
	if x.r != '/' {
		panic("incorrect comment start")
	}

	x.next()

	if x.r == '/' {
		return lexLineComment(x)
	} else if x.r == '*' {
		return lexBlockComment(x)
	}
	x.err("illegal char %q", x.r)
	return x.token(Illegal)
}

func lexLineComment(x *Lexer) *Token {
	for {
		x.next()
		if x.eof() || x.r == '\n' {
			break
		}
	}
	return x.token(Comment)
}

func lexBlockComment(x *Lexer) *Token {
	star := false
	for {
		x.next()
		if x.eof() {
			x.err("unexpected eof in block comment")
			return x.token(Comment)
		}

		if star && x.r == '/' {
			x.next()
			break
		}

		star = x.r == '*'
	}

	return x.token(Comment)
}
