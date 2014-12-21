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

func (vm *VirtMemory) translateRead(addr uint32, ring byte) (uint32, *Excep) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateRead(addr, ring)
}

func (vm *VirtMemory) translateWrite(addr uint32, ring byte) (uint32, *Excep) {
	if vm.ptable == nil {
		return addr, nil
	}
	return vm.ptable.TranslateWrite(addr, ring)
}

// ReadWordRing reads the byte at the given virtual address.
func (vm *VirtMemory) ReadWordRing(addr uint32, ring byte) (uint32, *Excep) {
	addr, e := vm.translateRead(addr, ring)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadWord(addr)
}

// WriteWordRing writes the byte at the given virtual address.
func (vm *VirtMemory) WriteWordRing(addr uint32, v uint32, ring byte) *Excep {
	addr, e := vm.translateWrite(addr, ring)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteWord(addr, v)
}

// ReadByteRing reads the byte at the given virtual address under a certain
// ring.
func (vm *VirtMemory) ReadByteRing(addr uint32, ring byte) (byte, *Excep) {
	addr, e := vm.translateRead(addr, ring)
	if e != nil {
		return 0, e
	}
	return vm.phyMem.ReadByte(addr)
}

// WriteByteRing writes a byte at the given virtual address under
// a certain ring.
func (vm *VirtMemory) WriteByteRing(addr uint32, v byte, ring byte) *Excep {
	addr, e := vm.translateWrite(addr, ring)
	if e != nil {
		return e
	}
	return vm.phyMem.WriteByte(addr, v)
}

// ReadWord reads the word at the given virtual address.
func (vm *VirtMemory) ReadWord(addr uint32) (uint32, *Excep) {
	return vm.ReadWordRing(addr, 0)
}

// WriteWord writes the word at the given virtual address.
func (vm *VirtMemory) WriteWord(addr uint32, v uint32) *Excep {
	return vm.WriteWordRing(addr, v, 0)
}

// ReadByte reads the byte at the given virtual address.
func (vm *VirtMemory) ReadByte(addr uint32) (byte, *Excep) {
	return vm.ReadByteRing(addr, 0)
}

// WriteByte writes a byte at the given virtual address.
func (vm *VirtMemory) WriteByte(addr uint32, v byte) *Excep {
	return vm.WriteByteRing(addr, v, 0)
}
