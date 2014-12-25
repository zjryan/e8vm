package arch8

import (
	"testing"
)

func TestPageTable(t *testing.T) {
	ne := func(e *Excep) {
		if e != nil {
			t.Fatal(e)
		}
	}

	as := func(cond bool) {
		if !cond {
			t.Fail()
		}
	}

	m := NewPhyMemory(4096 * PageSize) // 4 2nd level entries

	p8 := m.Page(8) // page eight
	p9 := m.Page(9) // page nine

	pte1 := ptEntry(0x9000)
	pte1.setBit(pteValid)
	p8.WriteWord(0, uint32(pte1))

	pte2 := ptEntry(0xa000)
	pte2.setBit(pteValid)
	for i := uint32(0); i < 512; i++ {
		// map all pages from 0 to 1023 to page 10
		p9.WriteWord(i*4, uint32(pte2))
	}

	pt := NewPageTable(m, 0x8000)

	for i := uint32(0); i < 512; i++ {
		ret, e := pt.Translate(i*0x1000+0x341, 0)
		ne(e)
		as(ret == 0xa341)
	}

	for i := uint32(512); i < 2048; i++ {
		ret, e := pt.Translate(i*0x1000, 0)
		as(ret == 0)
		as(e.Code == ErrPageFault)
	}

	addr := uint32(3*PageSize + 0x23)
	_, e := pt.TranslateRead(addr, 0)
	as(e == nil)

	pte2.setBit(pteReadonly)
	p9.WriteWord(3*4, uint32(pte2))
	_, e = pt.TranslateRead(addr, 0)
	as(e == nil)
	_, e = pt.TranslateWrite(addr, 0)
	as(e.Code == ErrPageReadonly)

	w := p9.ReadWord(3 * 4)
	as((w & (0x1 << pteUse)) != 0)
	as((w & (0x1 << pteDirty)) == 0)

	addr += PageSize
	_, e = pt.TranslateWrite(addr, 0)
	as(e == nil)
	w = p9.ReadWord(4 * 4)
	as((w & (0x1 << pteDirty)) != 0)
}
