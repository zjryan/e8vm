package parse

import (
	"lonnie.io/e8vm/lex8"
)

var regNameMap = map[string]uint32{
	"r0": 0,
	"r1": 1,
	"r2": 2,
	"r3": 3,
	"r4": 4,
	"r5": 5,
	"r6": 6,
	"r7": 7,

	"sp":  5,
	"ret": 6,
	"pc":  7,
}

func parseReg(p *parser, op *lex8.Token) uint32 {
	if op.Type != Operand {
		panic("not an operand")
	}

	ret, found := regNameMap[op.Lit]
	if !found {
		p.err(op.Pos, "invalid register name %q", op.Lit)
		return 0
	}
	return ret
}
