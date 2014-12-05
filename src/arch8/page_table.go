package arch8

// PageTable describes a page table in a physical memory.
// It can be used to translate a virtual memory address into a physical
// memory address.
type PageTable struct {
	mem  *PhyMemory // the physical memory
	root uint32     // root address

	// last translation
	pte1     ptEntry
	pte2     ptEntry
	pte1Addr uint32
	pte2Addr uint32
}

// NewPageTable creates a new page table pointer.
// The page table is saved at addr.
// If addr is not page size aligned, it is aligned down.
func NewPageTable(m *PhyMemory, addr uint32) *PageTable {
	addr -= addr % PageSize
	ret := new(PageTable)
	ret.mem = m
	ret.root = addr

	return ret
}

// bit [31:12] -> a page pointer
//
// bit 3: dirty bit
// bit 2: use bit
// bit 1: readonly bit
// bit 0: valid bit
type ptEntry uint32

const (
	pteValid    = 0
	pteReadonly = 1
	pteUse      = 2
	pteDirty    = 3
)

const u32one uint32 = 0x1

func (pte ptEntry) testBit(n uint) bool {
	return (uint32(pte) & (u32one << n)) != 0
}
func (pte *ptEntry) setBit(n uint) {
	*pte = ptEntry(uint32(*pte) | (u32one << n))
}
func (pte ptEntry) pn() uint32 {
	return uint32(pte / PageSize)
}

// Translate transalate a virutal address into physical address.
// It returns an error if the translation fails
func (pt *PageTable) Translate(addr uint32) (uint32, *Excep) {
	vpn := addr / PageSize
	off := addr % PageSize

	index1 := vpn / 1024
	index2 := vpn % 1024

	// first level page table entry
	pt.pte1Addr = pt.root + index1*4
	w, e := pt.mem.ReadWord(pt.pte1Addr)
	if e != nil {
		return 0, e
	}
	pt.pte1 = ptEntry(w)
	if !pt.pte1.testBit(pteValid) {
		return 0, newPageFault(addr)
	}
	pn1 := pt.pte1.pn()

	pt.pte2Addr = pn1*PageSize + index2*4
	w, e = pt.mem.ReadWord(pt.pte2Addr)
	if e != nil {
		return 0, e
	}
	pt.pte2 = ptEntry(w)
	if !pt.pte2.testBit(pteValid) {
		return 0, newPageFault(addr)
	}

	ppn := pt.pte2.pn()

	return ppn*PageSize + off, nil
}

func (pt *PageTable) updatePte() *Excep {
	e := pt.mem.WriteWord(pt.pte1Addr, uint32(pt.pte1))
	if e != nil {
		return e
	}

	e = pt.mem.WriteWord(pt.pte2Addr, uint32(pt.pte2))
	if e != nil {
		return e
	}

	return nil
}

// TranslateRead translates the address and sets the use bit.
func (pt *PageTable) TranslateRead(addr uint32) (uint32, *Excep) {
	ret, e := pt.Translate(addr)
	if e != nil {
		return 0, e
	}

	pt.pte1.setBit(pteUse)
	pt.pte2.setBit(pteUse)

	e = pt.updatePte()
	if e != nil {
		return 0, e
	}

	return ret, nil
}

// TranslateWrite translates the address and sets the use and dirty bit
func (pt *PageTable) TranslateWrite(addr uint32) (uint32, *Excep) {
	ret, e := pt.Translate(addr)
	if e != nil {
		return 0, e
	}

	if pt.pte1.testBit(pteReadonly) || pt.pte2.testBit(pteReadonly) {
		return 0, newPageReadonly(addr)
	}

	pt.pte1.setBit(pteUse)
	pt.pte1.setBit(pteDirty)
	pt.pte2.setBit(pteUse)
	pt.pte2.setBit(pteDirty)

	e = pt.updatePte()
	if e != nil {
		return 0, e
	}

	return ret, nil
}
