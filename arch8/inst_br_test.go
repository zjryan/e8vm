package arch8

import (
	"testing"
)

func TestInstBr(t *testing.T) {
	m := NewPhyMemory(PageSize * 16)
	cpu := NewCPU(m, new(InstBr), 0)

	for i := 0; i < 4; i++ {
		cpu.regs[i] = uint32(i)
	}

	bne := func(s1, s2 uint32, off int32) uint32 {
		in := uint32(32) << 24
		in |= (s1 & 0x7) << 21
		in |= (s2 & 0x7) << 18
		in |= uint32(off) & 0x3ffff
		return in
	}

	beq := func(s1, s2 uint32, off int32) uint32 {
		in := uint32(33) << 24
		in |= (s1 & 0x7) << 21
		in |= (s2 & 0x7) << 18
		in |= uint32(off) & 0x3ffff
		return in
	}

	wr := func(addr, w uint32) {
		e := m.WriteWord(addr, w)
		if e != nil {
			t.Fatal(e)
		}
	}

	wr(InitPC, beq(0, 0, 1))
	wr(InitPC+4, beq(3, 3, -1))
	wr(InitPC+8, bne(1, 2, -2))

	tick := func(exp uint32) {
		e := cpu.Tick()
		if e != nil {
			t.Fatal("unexpected exception", e)
		}

		pc := cpu.regs[PC]
		if pc != exp {
			t.Fatalf("expect pc 0x%08x, got 0x%08x", exp, pc)
		}
	}

	tick(InitPC + 8)
	tick(InitPC + 4)
	tick(InitPC + 4)
	tick(InitPC + 4)
}
