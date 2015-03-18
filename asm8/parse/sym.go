package parse

import (
	"strings"

	"lonnie.io/e8vm/lex8"
)

func isSymbol(sym string) bool {
	if sym == "" {
		return false
	}
	r := sym[0]
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func isIdent(id string) bool {
	if id == "" {
		return false
	}
	for i, r := range id {
		if r >= '0' && r <= '9' {
			if i > 0 {
				continue
			}
			return false
		}

		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r == '_' || r == ':' {
			continue
		}
		return false
	}
	return true
}

func parseSym(p *parser, t *lex8.Token) (pack, sym string) {
	if t.Type != Operand {
		panic("symbol not an operand")
	}

	sym = t.Lit
	dot := strings.Index(sym, ".")
	if dot >= 0 {
		pack, sym = sym[:dot], sym[dot+1:]
	}

	if dot >= 0 && !lex8.IsPkgName(pack) {
		p.Errorf(t.Pos, "invalid package name: %q", pack)
	} else if !isIdent(sym) {
		p.Errorf(t.Pos, "invalid symbol: %q", t.Lit)
	}

	return pack, sym
}
