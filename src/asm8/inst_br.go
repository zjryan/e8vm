package asm8

import (
	"lex8"
)

var (
	// op reg reg label
	opBrMap = map[string]uint32{
		"bne": 32,
		"beq": 33,
	}
)

func makeInstBr(op, s1, s2 uint32) *inst {
	ret := (op & 0xff) << 24
	ret |= (s1 & 0x7) << 21
	ret |= (s2 & 0x7) << 18
	return &inst{inst: ret}
}

func parseInstBr(p *Parser, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	var (
		op, s1, s2 uint32
		lab        string

		found bool
	)

	if op, found = opBrMap[opName]; found {
		// op reg reg label
		if argCount(p, ops, 3) {
			s1 = parseReg(p, args[0])
			s2 = parseReg(p, args[1])
			t := args[2]
			if parseLabel(p, t) {
				lab = t.Lit
			} else {
				p.err(t.Pos, "expects a label for %s", opName)
			}
		}
	} else {
		return nil, false
	}

	ret := makeInstBr(op, s1, s2)
	ret.symbol = lab
	ret.fill = fillLabel

	return ret, true
}
