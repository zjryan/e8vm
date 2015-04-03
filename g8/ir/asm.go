package ir

import (
	A "lonnie.io/e8vm/arch8"
	S "lonnie.io/e8vm/asm8"
)

// go lint (stupidly) forbids import with .
// so we will just copy the consts in here
const (
	_0   = A.R0
	_1   = A.R1
	_2   = A.R2
	_3   = A.R3
	_4   = A.R4
	_ret = A.RET
	_sp  = A.SP
	_pc  = A.PC
)

// an empty struct for a separate namespace
type _s struct{}

var asm _s

func (_s) ims(op, d, s uint32, im int32) uint32 {
	return S.InstImm(op, d, s, uint32(im))
}
func (_s) lw(d, s uint32, im int32) uint32 {
	return asm.ims(A.LW, d, s, im)
}
func (_s) sw(d, s uint32, im int32) uint32 {
	return asm.ims(A.SW, d, s, im)
}
func (_s) sb(d, s uint32, im int32) uint32 {
	return asm.ims(A.SB, d, s, im)
}
func (_s) lb(d, s uint32, im int32) uint32 {
	return asm.ims(A.LB, d, s, im)
}
func (_s) lbu(d, s uint32, im int32) uint32 {
	return asm.ims(A.LBU, d, s, im)
}
func (_s) addi(d, s uint32, im int32) uint32 {
	return asm.ims(A.ADDI, d, s, im)
}

func (_s) lui(d, im uint32) uint32     { return S.InstImm(A.LUI, d, 0, im) }
func (_s) ori(d, s, im uint32) uint32  { return S.InstImm(A.ORI, d, s, im) }
func (_s) xori(d, s, im uint32) uint32 { return S.InstImm(A.XORI, d, s, im) }
func (_s) andi(d, s, im uint32) uint32 { return S.InstImm(A.ANDI, d, s, im) }

func (_s) reg(op, d, s1, s2 uint32) uint32 {
	return S.InstReg(op, d, s1, s2, 0, 0)
}

func (_s) add(d, s1, s2 uint32) uint32  { return asm.reg(A.ADD, d, s1, s2) }
func (_s) sub(d, s1, s2 uint32) uint32  { return asm.reg(A.SUB, d, s1, s2) }
func (_s) mul(d, s1, s2 uint32) uint32  { return asm.reg(A.MUL, d, s1, s2) }
func (_s) div(d, s1, s2 uint32) uint32  { return asm.reg(A.DIV, d, s1, s2) }
func (_s) mod(d, s1, s2 uint32) uint32  { return asm.reg(A.MOD, d, s1, s2) }
func (_s) and(d, s1, s2 uint32) uint32  { return asm.reg(A.AND, d, s1, s2) }
func (_s) or(d, s1, s2 uint32) uint32   { return asm.reg(A.OR, d, s1, s2) }
func (_s) xor(d, s1, s2 uint32) uint32  { return asm.reg(A.XOR, d, s1, s2) }
func (_s) nor(d, s1, s2 uint32) uint32  { return asm.reg(A.NOR, d, s1, s2) }
func (_s) sltu(d, s1, s2 uint32) uint32 { return asm.reg(A.SLTU, d, s1, s2) }
func (_s) slt(d, s1, s2 uint32) uint32  { return asm.reg(A.SLT, d, s1, s2) }

func (_s) beq(s1, s2 uint32, im int32) uint32 {
	return S.InstBr(A.BEQ, s1, s2, im)
}
func (_s) bne(s1, s2 uint32, im int32) uint32 {
	return S.InstBr(A.BNE, s1, s2, im)
}

func (_s) jal(im int32) uint32 { return S.InstJmp(A.JAL, im) }
func (_s) j(im int32) uint32   { return S.InstJmp(A.J, im) }

func (_s) halt() uint32 { return S.InstSys(A.HALT, 0) }
