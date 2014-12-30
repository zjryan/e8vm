package arch8

import (
	"testing"
)

func TestVirtMemory(t *testing.T) {
	as := func(cond bool, s string, args ...interface{}) {
		if !cond {
			t.Fatalf(s, args...)
		}
	}

	eo := func(cond bool, s string, args ...interface{}) {
		if cond {
			t.Fatalf(s, args...)
		}
	}

	m := newPhyMemory(8 * PageSize)
	p1 := m.Page(1)
	p2 := m.Page(2)
	p3 := m.Page(3)

	pte1 := ptEntry(2 * PageSize)
	pte1.setBit(pteValid)

	p1.WriteWord(0, uint32(pte1))
	pte1 += PageSize
	p1.WriteWord(4*10, uint32(pte1))

	for i := uint32(0); i < 4; i++ {
		pte := ptEntry((i + 4) * PageSize)
		pte.setBit(pteValid)
		p2.WriteWord(4*i, uint32(pte))
		p3.WriteWord(4*(3-i), uint32(pte))
	}

	vm := newVirtMemory(m)
	vm.SetTable(PageSize)

	// page 0, 1, 2, 3 is mapped to 4, 5, 6, 7
	// page B+0, B+1, B+2, B+3 is mapped to 7, 6, 5, 4
	// where B = 10 * 1024
	w := uint32(0x37215a40)
	off := uint32(324)

	for i := uint32(0); i < 4; i++ {
		wd := w + i
		e := vm.WriteWord(off+i*PageSize, 0, wd)
		as(e == nil, "write error: %s", e)
		w2, e := vm.ReadWord((10*1024+(3-i))*PageSize+off, 0)
		as(e == nil, "read error: %s", e)
		eo(w2 != wd, "expect 0x%08x, got 0x%08x", wd, w2)
	}

	b := byte(0x54)
	off = 575
	for i := uint32(0); i < 4; i++ {
		bt := b + byte(i)
		e := vm.WriteByte(off+i*PageSize, 0, bt)
		as(e == nil, "write error: %s", e)
		b2, e := vm.ReadByte((10*1024+(3-i))*PageSize+off, 0)
		as(e == nil, "read error: %s", e)
		eo(b2 != bt, "expect 0x%02x, got 0x%02x", bt, b2)
	}
}
