package asm8

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/lex8"
)

func parseDataStr(p *parser, args []*lex8.Token) ([]byte, uint32) {
	buf := new(bytes.Buffer)

	for _, arg := range args {
		if arg.Type != String {
			p.err(arg.Pos, "expect string, got %s", typeStr(arg.Type))
			return nil, 0
		}

		s, e := strconv.Unquote(arg.Lit)
		if e != nil {
			p.err(arg.Pos, "invalid string %s", arg.Lit)
			return nil, 0
		}
		buf.Write([]byte(s))
	}

	return buf.Bytes(), 0
}
