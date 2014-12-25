package asm8

import (
	"lonnie.io/e8vm/lex8"
)

var opJmpMap = map[string]uint32{
	"j":   2,
	"jal": 3,
}

func isValidSymbol(sym string) bool {
	return true
}

func parseInstJmp(p *Parser, ops []*lex8.Token) (*inst, bool) {
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

	var pack, sym string
	var fill int
	var symTok *lex8.Token

	if argCount(p, ops, 1) {
		if parseLabel(p, ops[1]) {
			sym = ops[1].Lit
			symTok = ops[1]
			fill = fillLabel
		} else {
			pack, sym = parseSym(p, ops[1])
			fill = fillLink
		}
	}

	ret := &inst{
		inst:   (op & 0x3) << 30,
		pack:   pack,
		symbol: sym,
		fill:   fill,
		symTok: symTok,
	}
	return ret, true
}
