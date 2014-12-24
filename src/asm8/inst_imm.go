package asm8

import (
	"strconv"

	"lex8"
)

var (
	// op reg reg imm(signed)
	opImsMap = map[string]uint32{
		"addi": 1,
		"slti": 2,

		"lw":  6,
		"lb":  7,
		"lbu": 8,
		"sw":  9,
		"sb":  10,
	}

	// op reg reg imm(unsigned)
	opImuMap = map[string]uint32{
		"andi": 3,
		"ori":  4,
	}

	// op reg imm(signed or unsigned)
	opImmMap = map[string]uint32{
		"lui": 5,
	}
)

// parseImu parses an unsigned 16-bit immediate
func parseImu(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseUint(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid unsigned immediate %q: %s", op.Lit, e)
		return 0
	}

	if (ret & 0xffff) != ret {
		p.err(op.Pos, "immediate too large: %s", op.Lit)
		return 0
	}

	return uint32(ret)
}

// parseIms parses an unsigned 16-bit immediate
func parseIms(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid signed immediate %q: %s", op.Lit, e)
		return 0
	}

	if ret > 0x7fff || ret < -0x8000 {
		p.err(op.Pos, "immediate out of 16-bit range: %s", op.Lit)
		return 0
	}

	return uint32(ret) & 0xffff
}

// parseImm parses an unsigned 16-bit immediate
func parseImm(p *Parser, op *lex8.Token) uint32 {
	ret, e := strconv.ParseInt(op.Lit, 0, 32)
	if e != nil {
		p.err(op.Pos, "invalid signed immediate %q: %s", op.Lit, e)
		return 0
	}

	if ret > 0xffff || ret < -0x8000 {
		p.err(op.Pos, "immediate out of 16-bit range: %s", op.Lit)
		return 0
	}

	return uint32(ret) & 0xffff
}

func makeInstImm(op, d, s, im uint32) *inst {
	ret := uint32(0)
	ret |= (op & 0xff) << 24
	ret |= (d & 0x7) << 21
	ret |= (s & 0x7) << 18
	ret |= (im & 0xffff)

	return &inst{inst: ret}
}

func parseInstImm(p *Parser, ops []*lex8.Token) (*inst, bool) {
	op0 := ops[0]
	opName := op0.Lit
	args := ops[1:]

	var (
		op, d, s, im uint32
		pack, sym    string
		fill         int
	)

	argCount := func(n int) bool {
		if !argCount(p, ops, n) {
			return false
		}
		if n >= 1 {
			d = parseReg(p, args[0])
		}
		return true
	}

	parseSym := func(t *lex8.Token, f func(*Parser, *lex8.Token) uint32) {
		if isSymbol(t.Lit) {
			pack, sym = parseSym(p, t)
			fill = fillLow
		} else {
			im = f(p, t)
		}
	}

	var found bool
	if op, found = opImsMap[opName]; found {
		// op reg reg imm(signed)
		if argCount(3) {
			s = parseReg(p, args[1])
			parseSym(args[2], parseIms)
		}
	} else if op, found = opImuMap[opName]; found {
		// op reg reg imm(unsigned)
		if argCount(3) {
			s = parseReg(p, args[1])
			parseSym(args[2], parseImu)
		}
	} else if op, found = opImmMap[opName]; found {
		// op reg imm(signed or unsigned)
		if argCount(2) {
			parseSym(args[1], parseImm)
		}
		if opName == "lui" {
			fill = fillHigh
		}
	} else {
		return nil, false
	}

	ret := makeInstImm(op, d, s, im)
	ret.pack = pack
	ret.symbol = sym
	ret.fill = fill

	return ret, true
}
