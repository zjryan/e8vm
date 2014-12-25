package arch8

// PhyMemory is a collection of contiguous pages.
type PhyMemory struct {
	npage uint32
	pages map[uint32]*Page
}

const (
	pageVoid      = 0
	pageInterrupt = 1
	pageBasicIO   = 2

	pageBootImage = 8
)

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

// Size returns the size of the physical memory.
func (pm *PhyMemory) Size() uint32 {
	return pm.npage * PageSize
}

// Page returns the page for the particular page number
// Returns nil when the page number is out of range
func (pm *PhyMemory) Page(pn uint32) *Page {
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

func (pm *PhyMemory) pageForByte(addr uint32) (*Page, *Excep) {
	p := pm.Page(addr / PageSize)
	if p == nil {
		return nil, errOutOfRange
	}
	return p, nil
}

func (pm *PhyMemory) pageForWord(addr uint32) (*Page, *Excep) {
	if addr%4 != 0 {
		return nil, errMisalign
	}
	return pm.pageForByte(addr)
}

// ReadByte reads the byte at the given address.
// If the address is out of range, it returns an error.
func (pm *PhyMemory) ReadByte(addr uint32) (byte, *Excep) {
	p, e := pm.pageForByte(addr)
	if e != nil {
		return 0, e
	}
	return p.ReadByte(addr), nil
}

// WriteByte writes the byte at the given address.
// If the address is out of range, it returns an error.
func (pm *PhyMemory) WriteByte(addr uint32, v byte) *Excep {
	p, e := pm.pageForByte(addr)
	if e != nil {
		return e
	}
	p.WriteByte(addr, v)
	return e
}

// ReadWord reads the byte at the given address.
// If the address is out of range or not 4-byte aligned, it returns an error.
func (pm *PhyMemory) ReadWord(addr uint32) (uint32, *Excep) {
	p, e := pm.pageForWord(addr)
	if e != nil {
		return 0, e
	}
	return p.ReadWord(addr), nil
}

// WriteWord reads the byte at the given address.
// If the address is out of range or not 4-byte aligned, it returns an error.
func (pm *PhyMemory) WriteWord(addr uint32, v uint32) *Excep {
	p, e := pm.pageForWord(addr)
	if e != nil {
		return e
	}
	p.WriteWord(addr, v)
	return nil
}
