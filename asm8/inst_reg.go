package asm8

import (
	"strconv"

	"lonnie.io/e8vm/lex8"
)

var (
	// op reg reg shift
	opShiftMap = map[string]uint32{
		"sll": 0,
		"srl": 1,
		"sra": 2,
	}

	// op reg reg reg
	opReg3Map = map[string]uint32{
		"sllv": 3,
		"srlv": 4,
		"srla": 5,
		"add":  6,
		"sub":  7,
		"and":  8,
		"or":   9,
		"xor":  10,
		"nor":  11,
		"slt":  12,
		"sltu": 13,
		"mul":  14,
		"mulu": 15,
		"div":  16,
		"divu": 17,
		"mod":  18,
		"modu": 19,
	}

	// op reg reg
	opReg2Map = map[string]uint32{
		"mov": 0,
	}

	// op reg reg reg
	opFloatMap = map[string]uint32{
		"fadd": 0,
		"fsub": 1,
		"fmul": 2,
		"fdiv": 3,
		"fint": 4,
	}
)

func parseShift(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseUint(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid shift %q: %s", op.Lit, e)
		return 0
	}

	if (ret & 0x1f) != ret {
		p.err(op.Pos, "shift too large: %s", op.Lit)
		return 0
	}

	return uint32(ret)
}

func makeInstReg(fn, d, s1, s2, sh, isFloat uint32) *inst {
	ret := uint32(0)
	ret |= (d & 0x7) << 21
	ret |= (s1 & 0x7) << 18
	ret |= (s2 & 0x7) << 15
	ret |= (sh & 0x1f) << 10
	ret |= (isFloat & 0x1) << 8
	ret |= fn & 0xff

	return &inst{inst: ret}
}

func parseInstReg(p *Parser, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	// common args
	var fn, d, s1, s2, sh, isFloat uint32

	argCount := func(n int) bool {
		if !argCount(p, ops, n) {
			return false
		}
		if n >= 2 {
			d = parseReg(p, args[0])
			s1 = parseReg(p, args[1])
		}
		return true
	}

	var found bool
	if fn, found = opShiftMap[opName]; found {
		// op reg reg shift
		if argCount(3) {
			sh = parseShift(p, args[2])
		}
	} else if fn, found = opReg3Map[opName]; found {
		// op reg reg reg
		if argCount(3) {
			s2 = parseReg(p, args[2])
		}
	} else if fn, found = opReg2Map[opName]; found {
		// op reg reg
		argCount(2)
	} else if fn, found = opFloatMap[opName]; found {
		// op reg reg reg
		if argCount(3) {
			s2 = parseReg(p, args[2])
		}
		isFloat = 1
	} else {
		return nil, false
	}

	return makeInstReg(fn, d, s1, s2, sh, isFloat), true
}
