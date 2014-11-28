package arch8

import (
	"testing"

	"math/rand"
)

func TestInstReg(t *testing.T) {
	m := NewPhyMemory(PageSize * 32)
	cpu := NewCPU(m, new(InstReg), 0)

	tst := func(op, s1, s2, v1, v2, d, res uint32) {
		cpu.Reset()

		in := op & 0xff
		in |= (s1 & 0x7) << 18
		in |= (s2 & 0x7) << 15
		in |= (d & 0x7) << 21
		m.WriteWord(InitPC, in)

		cpu.regs[s1] = v1
		cpu.regs[s2] = v2
		e := cpu.Tick()
		if e != nil {
			t.Fatal("unexpected exception")
		}

		got := cpu.regs[d]
		if got != res {
			t.Fatalf("got 0x%08x, expect 0x%08x", got, res)
		}
	}

	tf := func(op uint32, f func(a, b uint32) uint32) {
		for i := 0; i < 100; i++ {
			s1 := uint32(rand.Intn(5))
			s2 := uint32(rand.Intn(5))

			for s2 == s1 {
				s2 = uint32(rand.Intn(5))
			}

			d := uint32(rand.Intn(5))
			v1 := uint32(rand.Int63())
			v2 := uint32(rand.Int63())

			exp := f(v1, v2)
			tst(op, s1, s2, v1, v2, d, exp)
		}
	}

	tst(0, 0, 0, 0, 0, 0, 0)

	tf(6, func(a, b uint32) uint32 { return a + b })
	tf(7, func(a, b uint32) uint32 { return a - b })
	tf(8, func(a, b uint32) uint32 { return a & b })
	tf(9, func(a, b uint32) uint32 { return a | b })
	tf(10, func(a, b uint32) uint32 { return a ^ b })
	tf(11, func(a, b uint32) uint32 { return ^(a | b) })
	tf(12, func(a, b uint32) uint32 {
		if int32(a) < int32(b) {
			return 1
		}
		return 0
	})
	tf(13, func(a, b uint32) uint32 {
		if a < b {
			return 1
		}
		return 0
	})
}
