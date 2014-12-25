package dasm8

import (
	"fmt"

	"encoding/binary"
)

// Dasm disassembles a byte block.
func Dasm(bs []byte, addr uint32) []*Line {
	var ret []*Line

	base := addr

	nline := len(bs) / 4
	for i := 0; i < nline; i++ {
		off := i * 4
		inst := binary.LittleEndian.Uint32(bs[off : off+4])
		ret = append(ret, NewLine(addr, inst))
		addr += 4
	}

	// link the jumps
	for _, line := range ret {
		if !line.IsJump {
			continue
		}

		index := int(line.To-base) / 4
		if index < nline {
			line.ToLine = ret[index]
		}
	}

	return ret
}

// Print prints a byte block as assembly lines.
func Print(bs []byte, addr uint32) {
	lines := Dasm(bs, addr)

	for _, line := range lines {
		fmt.Println(line)
	}
}
