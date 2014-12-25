package dasm8

import (
	"fmt"
)

// Line is a disassembled line
type Line struct {
	Addr uint32

	Bytes []byte
	Inst  uint32
	Str   string

	IsJump     bool
	Target     uint32
	TargetLine *Line
}

func (line *Line) String() string {
	if !line.IsJump {
		return fmt.Sprintf("%08x:  %08x   %s",
			line.Addr,
			line.Inst,
			line.Str,
		)
	}

	return fmt.Sprintf("%08x:  %08x   %s (%08x)",
		line.Addr,
		line.Inst,
		line.Str,
		line.Target,
	)
}
