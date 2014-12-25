package dasm8

import (
	"fmt"
)

// Dasm disassembles a byte block.
func Dasm(bs []byte, addr uint32) []*Line {
	var ret []*Line

	return ret
}

// Print prints a byte block as assembly lines.
func Print(bs []byte, addr uint32) {
	lines := Dasm(bs, addr)

	for _, line := range lines {
		fmt.Println(line)
	}
}
