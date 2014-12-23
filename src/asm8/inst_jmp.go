package asm8

import (
	"lex8"
)

var opJmpMap = map[string]uint32{
	"j":   2,
	"jal": 3,
}

func isValidSymbol(sym string) bool {
	return true
}

func parseInstJmp(p *Parser, ops []*lex8.Token) (*inst, bool) {
	if len(ops) == 0 {
		panic("0 ops")
	}

	op0 := ops[0]
	opName := op0.Lit
	var op uint32

	// op sym
	switch opName {
	case "j":
		op = 2
	case "jal":
		op = 3
	default:
		return nil, false
	}

	var sym string
	var fill int

	if argCount(p, ops, 1) {
		sym = ops[1].Lit
		if isLabel(sym) {
			if !isValidLabel(sym) {
				p.err(ops[1].Pos, "invalid label: %q", sym)
			}
			fill = fillLabel
		} else {
			if !isValidSymbol(sym) {
				p.err(ops[1].Pos, "invalid symbol: %q", sym)
			}
			fill = fillLink
		}
	}

	ret := &inst{
		inst:   (op & 0x3) << 30,
		symbol: sym,
		fill:   fill,
	}
	return ret, true
}
