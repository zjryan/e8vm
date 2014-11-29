package arch8

import (
	"testing"

	"math/rand"
)

func TestInstImm(t *testing.T) {
	m := NewPhyMemory(PageSize * 32)
	cpu := NewCPU(m, new(InstImm), 0)

	tst := func(op, s, v, im, d, res uint32) {
		cpu.Reset()

		in := (op & 0xff) << 24
		in |= (s & 0x7) << 18
		in |= (d & 0x7) << 21
		in |= im & 0xffff

		m.WriteWord(InitPC, in)

		cpu.regs[s] = v
		e := cpu.Tick()
		if e != nil {
			t.Fatal("unexpected exception")
		}

		got := cpu.regs[d]
		if got != res {
			t.Fatalf("got 0x%08x, expect 0x%08x", got, res)
		}
	}

	tf := func(op uint32, f func(v, im uint32) uint32) {
		for i := 0; i < 100; i++ {
			s := uint32(rand.Intn(5))
			d := uint32(rand.Intn(5))
			v := uint32(rand.Int63())
			im := uint32(rand.Int63()) & 0xffff

			exp := f(v, im)

			tst(op, s, v, im, d, exp)
		}
	}

	ti := func(op, v, im, exp uint32) {
		for s := uint32(0); s < 5; s++ {
			for d := uint32(0); d < 5; d++ {
				tst(op, s, v, im, d, exp)
			}
		}
	}

	sext := func(i uint32) uint32 {
		return uint32((int32(i << 16)) >> 16)
	}

	tf(1, func(v, im uint32) uint32 { return v + sext(im) })
	tf(2, func(v, im uint32) uint32 {
		se := sext(im)
		if int32(v) < int32(se) {
			return 1
		} else {
			return 0
		}
	})
	tf(3, func(v, im uint32) uint32 { return v & im })
	tf(4, func(v, im uint32) uint32 { return v | im })
	tf(5, func(_, im uint32) uint32 { return im << 16 })

	ti(6, 0, 0, 0)
}
