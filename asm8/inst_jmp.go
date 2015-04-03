package asm8

import (
	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/lex8"
)

var opJmpMap = map[string]uint32{
	"j":   2,
	"jal": 3,
}

func isValidSymbol(sym string) bool {
	return true
}

// InstJmp makes a jump instruction
func InstJmp(op uint32, im int32) uint32 {
	ret := (op & 0x3) << 30
	ret |= uint32(im) & 0x3fffffff
	return ret
}

func resolveInstJmp(p lex8.Logger, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	var op uint32

	// op sym
	switch opName {
	case "j":
		op = arch8.J
	case "jal":
		op = arch8.JAL
	default:
		return nil, false
	}

	var pack, sym string
	var fill int
	var symTok *lex8.Token

	if argCount(p, ops, 1) {
		symTok = ops[1]
		if checkLabel(p, ops[1]) {
			sym = ops[1].Lit
			fill = fillLabel
		} else {
			pack, sym = parseSym(p, ops[1])
			fill = fillLink
		}
	}

	ret := new(inst)
	ret.inst = InstJmp(op, 0)
	ret.pkg = pack
	ret.sym = sym
	ret.fill = fill
	ret.symTok = symTok

	return ret, true
}
