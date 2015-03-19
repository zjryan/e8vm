package asm8

import (
	"lonnie.io/e8vm/lex8"
)

var (
	// op
	opSysMap = map[string]uint32{
		"halt":    64,
		"syscall": 65,
		"usermod": 66,
		"iret":    68,
	}

	// op reg
	opSys1Map = map[string]uint32{
		"vtable": 67,
		"cpuid":  69,
	}
)

func makeInstSys(op, reg uint32) *inst {
	ret := ((op & 0xff) << 24) | ((reg & 0x7) << 21)
	return &inst{inst: ret}
}

func resolveInstSys(p lex8.Logger, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]
	var op, reg uint32

	argCount := func(n int) bool {
		if !argCount(p, ops, n) {
			return false
		}

		if n >= 1 {
			reg = resolveReg(p, args[0])
		}
		return true
	}

	var found bool
	if op, found = opSysMap[opName]; found {
		// op
		argCount(0)
	} else if op, found = opSys1Map[opName]; found {
		// op reg
		argCount(1)
	} else {
		return nil, false
	}

	return makeInstSys(op, reg), true
}
