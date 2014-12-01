package asm8

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

func lexOperand(x *Lexer) *Token {
	if !isOperandChar(x.r) {
		panic("incorrect operand start")
	}

	for {
		x.next()
		if x.eof() || !isOperandChar(x.r) {
			break
		}
	}

	ret := x.token(Operand)
	if isKeyword(ret.Lit) {
		ret.Type = Keyword
	}
	return ret
}
