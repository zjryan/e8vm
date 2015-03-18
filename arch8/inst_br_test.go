package arch8

import (
	"testing"
	
	"lonnie.io/e8vm/conv"
)


func TestInstBr(t *testing.T) {
	const initPC = conv.InitPC

	m := newPhyMemory(PageSize * 16)
	cpu := newCPU(m, new(instBr), 0)

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

	wr(initPC, beq(0, 0, 1))
	wr(initPC+4, beq(3, 3, -1))
	wr(initPC+8, bne(1, 2, -2))

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

	tick(initPC + 8)
	tick(initPC + 4)
	tick(initPC + 4)
	tick(initPC + 4)
}
