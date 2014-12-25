package dasm8

import (
	"fmt"
)

var (
	opJmpMap = map[uint32]string{
		2: "j",
		3: "jal",
	}
)

func instJmp(addr uint32, in uint32) *Line {
	op := (in >> 30) & 0x3
	off := int32((in & 0x3fffffff) << 2)
	ret := newLine(addr, in)

	if opStr, found := opJmpMap[op]; found {
		s := fmt.Sprintf("%s %d", opStr, off)

		ret.Str = s
		ret.IsJump = true
		ret.To = addr + 4 + uint32(off)
	}

	return ret
}
