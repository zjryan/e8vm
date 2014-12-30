package dasm8

import (
	"fmt"

	"encoding/binary"
)

// Dasm disassembles a byte block.
func Dasm(bs []byte, addr uint32) []*Line {
	var ret []*Line

	base := addr

	add := func(b []byte) {
		inst := binary.LittleEndian.Uint32(b)
		ret = append(ret, NewLine(addr, inst))
		addr += 4
	}

	nline := len(bs) / 4
	for i := 0; i < nline; i++ {
		off := i * 4
		add(bs[off : off+4])
	}

	// residue
	if len(bs)%4 != 0 {
		var buf [4]byte
		copy(buf[:], bs[nline*4:])
		add(buf[:])
	}

	// link the jumps
	for _, line := range ret {
		if !line.IsJump {
			continue
		}

		index := int(line.To-base) / 4
		if index >= 0 && index < nline {
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
