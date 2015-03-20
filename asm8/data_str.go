package asm8

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

func parseDataStr(p lex8.Logger, args []*lex8.Token) ([]byte, uint32) {
	if !checkTypeAll(p, args, parse.String) {
		return nil, 0
	}

	buf := new(bytes.Buffer)

	for _, arg := range args {
		if arg.Lit[0] != '"' {
			p.Errorf(arg.Pos, "expect string for string data")
			return nil, 0
		}

		s, e := strconv.Unquote(arg.Lit)
		if e != nil {
			p.Errorf(arg.Pos, "invalid string %s", arg.Lit)
			return nil, 0
		}
		buf.Write([]byte(s))
	}

	return buf.Bytes(), 0
}
