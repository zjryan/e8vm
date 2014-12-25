package dasm8

import (
	"fmt"
)

var (
	opBrMap = map[uint32]string{
		32: "bne",
		33: "beq",
	}
)

func instBr(addr uint32, in uint32) *Line {
	op := (in >> 24) & 0xff
	src1 := regStr((in >> 21) & 0x7)
	src2 := regStr((in >> 18) & 0x7)
	im := in & 0x3ffff
	off := int32(im<<14) >> 12

	ret := newLine(addr, in)
	if opStr, found := opBrMap[op]; found {
		s := fmt.Sprintf("%s %s %s %d", opStr, src1, src2, off)

		ret.Str = s
		ret.IsJump = true
		ret.To = addr + 4 + uint32(off)
	}

	return ret
}
