package dasm8

import (
	"fmt"

	"lonnie.io/e8vm/arch8"
)

var (
	opImsMap = map[uint32]string{
		arch8.ADDI: "addi",
		arch8.SLTI: "slti",
	}

	opMemMap = map[uint32]string{
		arch8.LW:  "lw",
		arch8.LB:  "lb",
		arch8.LBU: "lbu",
		arch8.SW:  "sw",
		arch8.SB:  "sb",
	}

	opImuMap = map[uint32]string{
		arch8.ANDI: "andi",
		arch8.ORI:  "ori",
		arch8.XORI: "xori",
	}

	opImu2Map = map[uint32]string{
		arch8.LUI: "lui",
	}
)

func instImm(addr uint32, in uint32) *Line {
	op := (in >> 24) & 0xff
	dest := regStr((in >> 21) & 0x7)
	src := regStr((in >> 18) & 0x7)
	imu := in & 0xffff
	ims := int32(imu<<16) >> 16

	var s string
	if opStr, found := opImsMap[op]; found {
		s = fmt.Sprintf("%s %s %s %d", opStr, dest, src, ims)
	} else if opStr, found := opMemMap[op]; found {
		if ims == 0 {
			s = fmt.Sprintf("%s %s %s", opStr, dest, src)
		} else {
			s = fmt.Sprintf("%s %s %s %d", opStr, dest, src, ims)
		}
	} else if opStr, found := opImuMap[op]; found {
		s = fmt.Sprintf("%s %s %s 0x%04x", opStr, dest, src, imu)
	} else if opStr, found := opImu2Map[op]; found {
		s = fmt.Sprintf("%s %s 0x%04x", opStr, dest, imu)
	}

	ret := newLine(addr, in)
	ret.Str = s

	return ret
}
