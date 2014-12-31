package dasm8

import (
	"bytes"
	"fmt"
)

// Line is a disassembled line
type Line struct {
	Addr uint32

	Inst uint32
	Str  string

	IsJump bool
	To     uint32
	ToLine *Line
}

func printables(inst uint32) string {
	ret := ""
	for i := 0; i < 4; i++ {
		r := rune(inst & 0xff)
		inst = inst >> 8
		if r >= ' ' && r <= '~' {
			ret = string(r) + ret
		} else {
			ret = "." + ret
		}
	}
	return ret
}

func (line *Line) String() string {
	ret := new(bytes.Buffer)
	fmt.Fprintf(ret, "%08x:  %08x", line.Addr, line.Inst)
	fmt.Fprintf(ret, "   %s", printables(line.Inst))
	fmt.Fprintf(ret, "    %s", line.Str)
	if line.IsJump {
		fmt.Fprintf(ret, "   // %08x", line.To)
	}

	return ret.String()
}

func newLine(addr uint32, in uint32) *Line {
	ret := new(Line)
	ret.Addr = addr
	ret.Inst = in
	return ret
}
