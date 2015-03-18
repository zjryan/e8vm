package parse

import (
	"lonnie.io/e8vm/lex8"
)

func isLabelStart(s string) bool {
	return len(s) > 0 && s[0] == '.'
}

func isLabel(s string) bool {
	if len(s) <= 1 || s[0] != '.' {
		return false
	}

	for i, r := range s[1:] {
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r == '_' {
			continue
		}
		return false
	}
	return true
}

func parseLabel(p *parser, t *lex8.Token) bool {
	if t.Type != Operand {
		panic("not an operand")
	}

	lab := t.Lit
	if !isLabelStart(lab) {
		return false
	}

	if !isLabel(lab) {
		p.Errorf(t.Pos, "invalid label: %q", lab)
	}

	return true
}
