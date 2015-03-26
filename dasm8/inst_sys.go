package dasm8

import (
	"fmt"

	"lonnie.io/e8vm/arch8"
)

var (
	opSysMap = map[uint32]string{
		arch8.HALT:    "halt",
		arch8.SYSCALL: "syscall",
		arch8.USERMOD: "usermod",
		arch8.IRET:    "iret",
	}

	opSys1Map = map[uint32]string{
		arch8.VTABLE: "vtable",
		arch8.CPUID:  "cpuid",
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
