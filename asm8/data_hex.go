package asm8

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
)

func parseDataHex(p lex8.Logger, args []*lex8.Token) ([]byte, uint32) {
	if !checkTypeAll(p, args, parse.Operand) {
		return nil, 0
	}

	buf := new(bytes.Buffer)
	for _, arg := range args {
		b, e := strconv.ParseUint(arg.Lit, 16, 8)
		if e != nil {
			p.Errorf(arg.Pos, "%s", e)
			return nil, 0
		}

		buf.WriteByte(byte(b))
	}

	return buf.Bytes(), 0
}
