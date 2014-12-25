package dasm8

import (
	"fmt"
)

var (
	opShiftMap = map[uint32]string{
		0: "sll",
		1: "srl",
		2: "sra",
	}

	opReg3Map = map[uint32]string{
		3:  "sllv",
		4:  "srlv",
		5:  "srla",
		6:  "add",
		7:  "sub",
		8:  "and",
		9:  "or",
		10: "xor",
		11: "nor",
		12: "slt",
		13: "sltu",
		14: "mul",
		15: "mulu",
		16: "div",
		17: "divu",
		18: "mod",
		19: "modu",
	}

	opFloatMap = map[uint32]string{
		0: "fadd",
		1: "fsub",
		2: "fmul",
		3: "fdiv",
		4: "fint",
	}
)

func instReg(addr uint32, in uint32) *Line {
	if ((in >> 24) & 0xff) != 0 {
		panic("not a register inst")
	}

	dest := regStr((in >> 21) & 0x7)
	src1 := regStr((in >> 18) & 0x7)
	src2 := regStr((in >> 15) & 0x7)
	shift := (in >> 10) & 0x1f
	isFloat := (in >> 8) & 0x1
	funct := in & 0xff

	var s string
	if isFloat == 0 {
		if funct == 0 && shift == 0 {
			s = fmt.Sprintf("mov %s %s", dest, src1)
		} else if opStr, found := opShiftMap[funct]; found {
			s = fmt.Sprintf("%s %s %s %d", opStr, dest, src1, shift)
		} else if opStr, found := opReg3Map[funct]; found {
			s = fmt.Sprintf("%s %s %s %s", opStr, dest, src1, src2)
		}
	} else {
		if opStr, found := opFloatMap[funct]; found {
			s = fmt.Sprintf("%s %s %s %s", opStr, dest, src1, src2)
		}
	}

	ret := newLine(addr, in)
	ret.Str = s

	return ret
}
