package dasm8

import (
	"fmt"
)

var (
	opSysMap = map[uint32]string{
		64: "halt",
		65: "syscall",
		66: "usermod",
		68: "iret",
	}

	opSys1Map = map[uint32]string{
		67: "vtable",
		69: "cpuid",
	}
)

func instSys(addr uint32, in uint32) *Line {
	op := (in >> 24) & 0xff
	src := regStr((in >> 21) & 0x7)

	var s string
	if opStr, found := opSysMap[op]; found {
		s = opStr
	} else if opStr, found := opSys1Map[op]; found {
		s = fmt.Sprintf("%s %s", opStr, src)
	}

	ret := newLine(addr, in)
	ret.Str = s
	return ret
}
