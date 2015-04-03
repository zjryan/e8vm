package asm8

import (
	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/lex8"
)

var (
	// op reg reg label
	opBrMap = map[string]uint32{
		"bne": arch8.BNE,
		"beq": arch8.BEQ,
	}
)

// InstBr compose a branch instruction
func InstBr(op, s1, s2 uint32, im int32) uint32 {
	ret := (op & 0xff) << 24
	ret |= (s1 & 0x7) << 21
	ret |= (s2 & 0x7) << 18
	ret |= uint32(im) & 0x3ffff
	return ret
}

func makeInstBr(op, s1, s2 uint32) *inst {
	ret := InstBr(op, s1, s2, 0)
	return &inst{inst: ret}
}

func resolveInstBr(p lex8.Logger, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	var (
		op, s1, s2 uint32
		lab        string
		symTok     *lex8.Token

		found bool
	)

	if op, found = opBrMap[opName]; found {
		// op reg reg label
		if argCount(p, ops, 3) {
			s1 = resolveReg(p, args[0])
			s2 = resolveReg(p, args[1])
			symTok = args[2]
			if checkLabel(p, symTok) {
				lab = symTok.Lit
			} else {
				p.Errorf(symTok.Pos, "expects a label for %s", opName)
			}
		}
	} else {
		return nil, false
	}

	ret := makeInstBr(op, s1, s2)
	ret.sym = lab
	ret.fill = fillLabel
	ret.symTok = symTok

	return ret, true
}
