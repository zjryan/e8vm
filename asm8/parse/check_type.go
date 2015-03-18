package parse

import (
	"lonnie.io/e8vm/lex8"
)

func checkAllType(p *parser, args []*lex8.Token, t int) bool {
	for _, arg := range args {
		if arg.Type != t {
			p.Errorf(arg.Pos, "expect %s, got %s", p.TypeStr(t), p.TypeStr(arg.Type))
			return false
		}
	}
	return true
}
