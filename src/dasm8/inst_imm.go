package dasm8

import (
	"fmt"
)

var (
	opImsMap = map[uint32]string{
		1: "addi",
		2: "slti",

		6:  "lw",
		7:  "lb",
		8:  "lbu",
		9:  "sw",
		10: "sb",
	}

	opImuMap = map[uint32]string{
		3: "andi",
		4: "ori",
		5: "lui",
	}
)

func instImm(addr uint32, in uint32) *Line {
	op := (in >> 24) & 0xff
	dest := regStr((in >> 21) & 0x7)
	src := regStr((in >> 18) & 0x7)
	imu := in & 0xffff
	ims := uint32(int32(imu<<16) >> 16)

	var s string
	if opStr, found := opImsMap[op]; found {
		s = fmt.Sprintf("%s %s %s %d", opStr, dest, src, ims)
	} else if opStr, found := opImuMap[op]; found {
		s = fmt.Sprintf("%s %s %s 0x%04x", opStr, dest, src, imu)
	}

	ret := newLine(addr, in)
	ret.Str = s

	return ret
}