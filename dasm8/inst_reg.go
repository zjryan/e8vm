package dasm8

import (
	"fmt"

	"lonnie.io/e8vm/arch8"
)

var (
	opShiftMap = map[uint32]string{
		arch8.SLL: "sll",
		arch8.SRL: "srl",
		arch8.SRA: "sra",
	}

	opReg3Map = map[uint32]string{
		arch8.SLLV: "sllv",
		arch8.SRLV: "srlv",
		arch8.SRLA: "srla",
		arch8.ADD:  "add",
		arch8.SUB:  "sub",
		arch8.AND:  "and",
		arch8.OR:   "or",
		arch8.XOR:  "xor",
		arch8.NOR:  "nor",
		arch8.SLT:  "slt",
		arch8.SLTU: "sltu",
		arch8.MUL:  "mul",
		arch8.MULU: "mulu",
		arch8.DIV:  "div",
		arch8.DIVU: "divu",
		arch8.MOD:  "mod",
		arch8.MODU: "modu",
	}

	opFloatMap = map[uint32]string{
		arch8.FADD: "fadd",
		arch8.FSUB: "fsub",
		arch8.FMUL: "fmul",
		arch8.FDIV: "fdiv",
		arch8.FINT: "fint",
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
