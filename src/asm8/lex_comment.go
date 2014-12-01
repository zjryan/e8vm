package asm8

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
