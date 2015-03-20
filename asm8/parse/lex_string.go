package parse

import (
	"unicode"

	"lonnie.io/e8vm/lex8"
)

func digitVal(r rune) int {
	switch {
	case '0' <= r && r <= '9':
		return int(r - '0')
	case 'a' <= r && r <= 'f':
		return int(r - 'a' + 10)
	case 'A' <= r && r <= 'F':
		return int(r - 'A' + 10)
	}
	return 16
}

func lexEscape(x *lex8.Lexer, quote rune) bool {
	var n int
	var base, max uint32
	if x.Ended() {
		x.Errorf("escape not terminated")
		return false
	}
	r := x.Rune()
	switch r {
	case 'a', 'b', 'f', 'n', 'r', 't', 'v', '\\', quote:
		x.Next()
		return true
	case '0', '1', '2', '3', '4', '5', '6', '7':
		n, base, max = 3, 8, 255
	case 'x':
		x.Next()
		n, base, max = 2, 16, 255
	case 'u':
		x.Next()
		n, base, max = 4, 16, unicode.MaxRune
	case 'U':
		x.Next()
		n, base, max = 8, 16, unicode.MaxRune
	default:
		x.Errorf("unknown escape sequence")
		return false
	}

	var v uint32
	for i := 0; i < n; i++ {
		if x.Ended() {
			x.Errorf("escape not terminated")
			return false
		}

		r := x.Rune()
		d := uint32(digitVal(r))
		if d >= base {
			x.Errorf("illegal escape char %#U", r)
			return false
		}

		v *= base
		v += d

		x.Next()
	}

	if v > max || 0xD800 <= v && v < 0xE000 {
		x.Errorf("invalid unicode code point")
		return false
	}

	return true
}

func lexString(x *lex8.Lexer) *lex8.Token {
	if !x.See('"') {
		panic("incorrect string start")
	}

	x.Next()
	for {
		if x.Ended() {
			x.Errorf("unexpected eof in string")
			break
		} else if x.See('\n') {
			x.Errorf("unexpected endl in string")
			break
		}

		if x.See('"') {
			x.Next()
			break
		} else if x.See('\\') {
			x.Next()
			lexEscape(x, '"')
		} else {
			x.Next()
		}
	}

	return x.MakeToken(String)
}
