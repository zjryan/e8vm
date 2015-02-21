package asm8

import (
	"lonnie.io/e8vm/asm8/ast"
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

func makeInstSys(op, reg uint32) *ast.Inst {
	ret := ((op & 0xff) << 24) | ((reg & 0x7) << 21)
	return &ast.Inst{Inst: ret}
}

func parseInstSys(p *parser, ops []*lex8.Token) (*ast.Inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]
	var op, reg uint32

	argCount := func(n int) bool {
		if !argCount(p, ops, n) {
			return false
		}

		if n >= 1 {
			reg = parseReg(p, args[0])
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
