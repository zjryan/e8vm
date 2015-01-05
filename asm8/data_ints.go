package asm8

import (
	"bytes"
	"strconv"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/lex8"
)

func nbitAlign(nbit int) uint32 {
	if nbit == 8 {
		return 0
	} else if nbit == 32 {
		return 4
	} else {
		panic("invalid nbit")
	}
}

const (
	modeWord   = 0x1
	modeSigned = 0x2
)

func parseDataInts(p *parser, args []*lex8.Token, mode int) ([]byte, uint32) {
	if !checkAllType(p, args, Operand) {
		return nil, 0
	}

	var i int64
	nbit := 8
	if mode&modeWord != 0 {
		nbit = 32
	}
	var e error

	buf := new(bytes.Buffer)
	for _, arg := range args {
		if mode&modeSigned != 0 {
			i, e = strconv.ParseInt(arg.Lit, 0, nbit)
		} else {
			var ui uint64
			ui, e = strconv.ParseUint(arg.Lit, 0, nbit)
			i = int64(ui)
		}
		if e != nil {
			p.err(arg.Pos, "%s", e)
			return nil, 0
		}

		if nbit == 8 {
			buf.WriteByte(byte(i))
		} else if nbit == 32 {
			var bs [4]byte
			arch8.Endian.PutUint32(bs[:], uint32(i))
			buf.Write(bs[:])
		}
	}

	return buf.Bytes(), nbitAlign(nbit)
}
