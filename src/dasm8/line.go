package dasm8

import (
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

func (line *Line) String() string {
	if !line.IsJump {
		return fmt.Sprintf("%08x:  %08x   %s",
			line.Addr,
			line.Inst,
			line.Str,
		)
	}

	return fmt.Sprintf("%08x:  %08x   %s  // %08x",
		line.Addr,
		line.Inst,
		line.Str,
		line.To,
	)
}

func newLine(addr uint32, in uint32) *Line {
	ret := new(Line)
	ret.Addr = addr
	ret.Inst = in
	return ret
}
