package arch8

import (
	"encoding/binary"
)

// PageSize is the number of bytes a page contains.
const PageSize = 4096

// Page is a memory addressable area of PageSize bytes
type Page struct {
	bytes []byte
}

// NewPage creates a new empty page.
func NewPage() *Page {
	return &Page{
		bytes: make([]byte, PageSize),
	}
}

func checkRange(offset uint32) {
	if offset > PageSize {
		panic("offset out of range")
	}
}

func checkWordAlign(offset uint32) {
	if offset%4 != 0 {
		panic("offset not 4-byte aligned")
	}
}

// Read reads a byte at the particular offset.
// It panics when the offset is larger than PageSize
func (p *Page) ReadByte(offset uint32) byte {
	checkRange(offset)
	return p.bytes[offset]
}

// Write writes a byte into the page at a particular offset.
// It panics when the offset is larger than PageSize.
func (p *Page) WriteByte(offset uint32, b byte) {
	checkRange(offset)
	p.bytes[offset] = b
}

// The machines endian.
var Endian = binary.LittleEndian

// ReadWord reads the word at the particular offset.
// It panics when the offset is larger than PageSize,
// or when the offset is not 4-byte aligned.
func (p *Page) ReadWord(offset uint32) uint32 {
	checkRange(offset)
	checkWordAlign(offset)
	return Endian.Uint32(p.bytes[offset : offset+4])
}

// WriteWord writes the word at the particular offset.
// It panics when the offset is larger than PageSize,
// or when the offset is not 4-byte aligned.
func (p *Page) WriteWord(offset uint32, w uint32) {
	checkRange(offset)
	checkWordAlign(offset)
	Endian.PutUint32(p.bytes[offset:offset+4], w)
}
