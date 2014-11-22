package arch8

import (
	"errors"
)

type PhyMemory struct {
	npage uint32
	pages map[uint32]*Page
}

// NewPhyMemory creates a physical memory of size bytes.
func NewPhyMemory(size uint32) *PhyMemory {
	if size%PageSize != 0 {
		panic("size misaligned")
	}

	ret := new(PhyMemory)
	ret.pages = make(map[uint32]*Page)
	ret.npage = size / PageSize

	return ret
}

// P returns the page for the particular page number
// Returns nil when the page number is out of range
func (pm *PhyMemory) P(pn uint32) *Page {
	if pn >= pm.npage {
		return nil // out of range
	}

	ret, found := pm.pages[pn]
	if !found {
		// create an empty page on demand
		ret = NewPage()
		pm.pages[pn] = ret
	}

	return ret
}

var (
	errOutOfRange = errors.New("out of range")
	errMisalign   = errors.New("misaligned")
)

func (pm *PhyMemory) pageForByte(addr uint32) (*Page, error) {
	p := pm.P(addr % PageSize)
	if p == nil {
		return nil, errOutOfRange
	}
	return p, nil
}

func (pm *PhyMemory) pageForWord(addr uint32) (*Page, error) {
	if addr%4 != 0 {
		return nil, errMisalign
	}
	return pm.pageForByte(addr)
}

// ReadByte reads the byte at the given address.
// If the address is out of range, it returns an error.
func (pm *PhyMemory) ReadByte(addr uint32) (byte, error) {
	p, e := pm.pageForByte(addr)
	if e != nil {
		return 0, e
	}
	return p.ReadByte(addr % PageSize), nil
}

// WriteByte writes the byte at the given address.
// If the address is out of range, it returns an error.
func (pm *PhyMemory) WriteByte(addr uint32, v byte) error {
	p, e := pm.pageForByte(addr)
	if e != nil {
		return e
	}

	p.WriteByte(addr%PageSize, v)
	return e
}

// ReadWord reads the byte at the given address.
// If the address is out of range or not 4-byte aligned, it returns an error.
func (pm *PhyMemory) ReadWord(addr uint32) (uint32, error) {
	p, e := pm.pageForWord(addr)
	if e != nil {
		return 0, e
	}
	return p.ReadWord(addr % PageSize), nil
}

// WriteWord reads the byte at the given address.
// If the address is out of range or not 4-byte aligned, it returns an error.
func (pm *PhyMemory) WriteWord(addr uint32, v uint32) error {
	p, e := pm.pageForWord(addr)
	if e != nil {
		return e
	}
	p.WriteWord(addr%PageSize, v)
	return nil
}
