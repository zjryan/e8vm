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
// If the address is not page size aligned, it will be aligned down.
// If the address is 0, it will use direct mapping.
func (vm *VirtMemory) SetTable(root uint32) {
	if root == 0 {
		vm.ptable = nil
	} else {
		vm.ptable = NewPageTable(vm.phyMem, root)
	}
}

func (vm *VirtMemory) transRead(addr uint32, ring byte) (uint32, *Excep) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateRead(addr, ring)
}

func (vm *VirtMemory) transWrite(addr uint32, ring byte) (uint32, *Excep) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateWrite(addr, ring)
}

// ReadWord reads the byte at the given virtual address.
func (vm *VirtMemory) ReadWord(addr uint32, ring byte) (uint32, *Excep) {
	addr, e := vm.transRead(addr, ring)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadWord(addr)
}

// WriteWord writes the byte at the given virtual address.
func (vm *VirtMemory) WriteWord(addr uint32, ring byte, v uint32) *Excep {
	addr, e := vm.transWrite(addr, ring)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteWord(addr, v)
}

// ReadByte reads the byte at the given virtual address.
func (vm *VirtMemory) ReadByte(addr uint32, ring byte) (byte, *Excep) {
	addr, e := vm.transRead(addr, ring)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadByte(addr)
}

// WriteByte writes a byte at the given virtual address under
// a certain ring.
func (vm *VirtMemory) WriteByte(addr uint32, ring byte, v byte) *Excep {
	addr, e := vm.transWrite(addr, ring)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteByte(addr, v)
}
