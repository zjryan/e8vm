package arch8

// VirtMemory defines an address space which page mapping
// is defined by a page table.
// If no page table is defined, direct mapping is used
type VirtMemory struct {
	phyMem *PhyMemory
	ptable *PageTable
}

// NewVirtMemory creates a new virtual address space with no page table.
func NewVirtMemory(phy *PhyMemory) *VirtMemory {
	ret := new(VirtMemory)
	ret.phyMem = phy
	return ret
}

// SetTable applies a particular pagetable at a physical memory position.
// If root is 0, direct mapping is used.
func (vm *VirtMemory) SetTable(root uint32) {
	if root == 0 {
		vm.ptable = nil
	} else {
		vm.ptable = &PageTable{
			mem:  vm.phyMem,
			root: root,
		}
	}
}

func (vm *VirtMemory) translateRead(addr uint32) (uint32, error) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateRead(addr)
}

func (vm *VirtMemory) translateWrite(addr uint32) (uint32, error) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateWrite(addr)
}

// ReadWord reads the byte at the given virtual address.
func (vm *VirtMemory) ReadWord(addr uint32) (uint32, error) {
	addr, e := vm.translateRead(addr)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadWord(addr)
}

// WriteWord writes the byte at the given virtual address.
func (vm *VirtMemory) WriteWord(addr uint32, v uint32) error {
	addr, e := vm.translateWrite(addr)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteWord(addr, v)
}

// ReadByte reads the byte at the given
func (vm *VirtMemory) ReadByte(addr uint32) (byte, error) {
	addr, e := vm.translateRead(addr)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadByte(addr)
}

// WriteByte writes the byte at the given
func (vm *VirtMemory) WriteByte(addr uint32, v byte) error {
	addr, e := vm.translateWrite(addr)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteByte(addr, v)
}
