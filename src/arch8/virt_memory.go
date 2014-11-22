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
