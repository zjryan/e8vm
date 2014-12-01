package asm8

func lexString(x *Lexer) *Token {
	escaped := false

	for {
		x.next()
		if x.eof() {
			x.err("unexpected eof in string")
			return x.token(String)
		} else if x.r == '\n' {
			x.err("unexpected endl in string")
			return x.token(String)
		}

		if escaped {
			escaped = false
		} else {
			if x.r == '\\' {
				escaped = true
			} else if x.r == '"' {
				x.next()
				break
			}
		}
	}

	return x.token(String)
}
