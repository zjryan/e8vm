package asm8

import (
	"strconv"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/lex8"
)

var (
	// op reg reg shift
	opShiftMap = map[string]uint32{
		"sll": arch8.SLL,
		"srl": arch8.SRL,
		"sra": arch8.SRA,
	}

	// op reg reg reg
	opReg3Map = map[string]uint32{
		"sllv": arch8.SLLV,
		"srlv": arch8.SRLV,
		"srla": arch8.SRLA,
		"add":  arch8.ADD,
		"sub":  arch8.SUB,
		"and":  arch8.AND,
		"or":   arch8.OR,
		"xor":  arch8.XOR,
		"nor":  arch8.NOR,
		"slt":  arch8.SLT,
		"sltu": arch8.SLTU,
		"mul":  arch8.MUL,
		"mulu": arch8.MULU,
		"div":  arch8.DIV,
		"divu": arch8.DIVU,
		"mod":  arch8.MOD,
		"modu": arch8.MODU,
	}

	// op reg reg
	opReg2Map = map[string]uint32{
		"mov": arch8.SLL,
	}

	// op reg reg reg
	opFloatMap = map[string]uint32{
		"fadd": arch8.FADD,
		"fsub": arch8.FSUB,
		"fmul": arch8.FMUL,
		"fdiv": arch8.FDIV,
		"fint": arch8.FINT,
	}
)

func parseShift(p lex8.Logger, op *lex8.Token) uint32 {
	ret, e := strconv.ParseUint(op.Lit, 0, 32)
	if e != nil {
		p.Errorf(op.Pos, "invalid shift %q: %s", op.Lit, e)
		return 0
	}

	if (ret & 0x1f) != ret {
		p.Errorf(op.Pos, "shift too large: %s", op.Lit)
		return 0
	}

	return uint32(ret)
}

// InstReg composes a register based instruction
func InstReg(fn, d, s1, s2, sh, isFloat uint32) uint32 {
	ret := uint32(0)
	ret |= (d & 0x7) << 21
	ret |= (s1 & 0x7) << 18
	ret |= (s2 & 0x7) << 15
	ret |= (sh & 0x1f) << 10
	ret |= (isFloat & 0x1) << 8
	ret |= fn & 0xff
	return ret
}

func makeInstReg(fn, d, s1, s2, sh, isFloat uint32) *inst {
	ret := InstReg(fn, d, s1, s2, sh, isFloat)
	return &inst{inst: ret}
}

func resolveInstReg(log lex8.Logger, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	// common args
	var fn, d, s1, s2, sh, isFloat uint32

	argCount := func(n int) bool {
		if !argCount(log, ops, n) {
			return false
		}
		if n >= 2 {
			d = resolveReg(log, args[0])
			s1 = resolveReg(log, args[1])
		}
		return true
	}

	var found bool
	if fn, found = opShiftMap[opName]; found {
		// op reg reg shift
		if argCount(3) {
			sh = parseShift(log, args[2])
		}
	} else if fn, found = opReg3Map[opName]; found {
		// op reg reg reg
		if argCount(3) {
			s2 = resolveReg(log, args[2])
		}
	} else if fn, found = opReg2Map[opName]; found {
		// op reg reg
		argCount(2)
	} else if fn, found = opFloatMap[opName]; found {
		// op reg reg reg
		if argCount(3) {
			s2 = resolveReg(log, args[2])
		}
		isFloat = 1
	} else {
		return nil, false
	}

	return makeInstReg(fn, d, s1, s2, sh, isFloat), true
}
