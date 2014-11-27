package arch8

import (
	"testing"
)

func TestVirtMemory(t *testing.T) {
	m := NewPhyMemory(8 * PageSize)
	p0 := m.P(0)
	p1 := m.P(1)
	p2 := m.P(2)

	pte1 := ptEntry(1 * PageSize)
	pte1.setBit(pteValid)
	p0.WriteWord(0, uint32(pte1))
	pte1 += PageSize
	p0.WriteWord(4*10, uint32(pte1))

	for i := uint32(0); i < 4; i++ {
		pte := ptEntry((i + 4) * PageSize)
		pte.setBit(pteValid)
		p1.WriteWord(4*i, uint32(pte))
		p2.WriteWord(4*(3-i), uint32(pte))
	}

}
