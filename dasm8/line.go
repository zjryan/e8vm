package dasm8

import (
	"bytes"
	"fmt"
	"encoding/binary"
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

func printables(bs []byte) string {
	ret := ""
	for _, b := range bs {
		r := rune(b)
		if r >= ' ' && r <= '~' {
			ret += string(r)
		} else {
			ret += "."
		}
	}
	return ret
}

var enc = binary.LittleEndian

func (line *Line) String() string {
	ret := new(bytes.Buffer)
	var buf [4]byte
	enc.PutUint32(buf[:], line.Inst)

	fmt.Fprintf(ret, "%08x:  % 02x", line.Addr, buf[:])
	fmt.Fprintf(ret, "   %s", printables(buf[:]))
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
