package arch8

import (
	"testing"

	"math"
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

	tsts := func(op, s1, sh, v1, d, res uint32) {
		cpu.Reset()

		in := op & 0xff
		in |= (sh & 0x1f) << 10
		in |= (s1 & 0x7) << 18
		in |= (d & 0x7) << 21
		m.WriteWord(InitPC, in)

		cpu.regs[s1] = v1
		e := cpu.Tick()
		if e != nil {
			t.Fatal("unexpected exception")
		}

		got := cpu.regs[d]
		if got != res {
			t.Fatalf("got 0x%08x, expect 0x%08x", got, res)
		}
	}

	tstf := func(op, s1, s2, d uint32, f1, f2, res float32) {
		cpu.Reset()

		in := op & 0xff
		in |= (s1 & 0x7) << 18
		in |= (s2 & 0x7) << 15
		in |= (d & 0x7) << 21
		in |= 0x1 << 8
		m.WriteWord(InitPC, in)

		cpu.regs[s1] = math.Float32bits(f1)
		cpu.regs[s2] = math.Float32bits(f2)
		e := cpu.Tick()
		if e != nil {
			t.Fatal("unexpected exception")
		}

		got := cpu.regs[d]
		fres := math.Float32bits(res)
		if got != fres {
			t.Fatalf("got 0x%08x, expect 0x%08x", got, fres)
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

	tf0 := func(op uint32, f func(a uint32) uint32) {
		for i := 0; i < 100; i++ {
			s1 := uint32(rand.Intn(5))
			s2 := uint32(rand.Intn(5))
			for s2 == s1 {
				s2 = uint32(rand.Intn(5))
			}

			d := uint32(rand.Intn(5))
			v1 := uint32(rand.Int63())
			exp := f(v1)
			tst(op, s1, s2, v1, 0, d, exp)
		}
	}

	tfs := func(op uint32, f func(a, b uint32) uint32) {
		for sh := uint32(0); sh < 32; sh++ {
			s1 := uint32(rand.Intn(5))

			d := uint32(rand.Intn(5))
			v1 := uint32(rand.Int63())

			exp := f(v1, sh)
			tsts(op, s1, sh, v1, d, exp)
		}
	}

	tff := func(op uint32, f func(a, b float32) float32) {
		for sh := uint32(0); sh < 32; sh++ {
			s1 := uint32(rand.Intn(5))
			s2 := uint32(rand.Intn(5))
			for s2 == s1 {
				s2 = uint32(rand.Intn(5))
			}
			d := uint32(rand.Intn(5))

			f1 := math.Float32frombits(uint32(rand.Int63()))
			f2 := math.Float32frombits(uint32(rand.Int63()))

			exp := f(f1, f2)
			tstf(op, s1, s2, d, f1, f2, exp)
		}
	}

	tst(0, 0, 0, 0, 0, 0, 0) // noop

	tf(3, func(a, b uint32) uint32 { return a << b })
	tf(4, func(a, b uint32) uint32 { return a >> b })
	tf(5, func(a, b uint32) uint32 { return uint32(int32(a) >> b) })

	tfs(0, func(a, sh uint32) uint32 { return a << sh })
	tfs(1, func(a, sh uint32) uint32 { return a >> sh })
	tfs(2, func(a, sh uint32) uint32 { return uint32(int32(a) >> sh) })

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
	tf(14, func(a, b uint32) uint32 {
		return uint32(int32(a) * int32(b))
	})
	tf(15, func(a, b uint32) uint32 { return a * b })

	tf(16, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return uint32(int32(a) / int32(b))
	})
	tf(17, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return a / b
	})
	tf0(16, func(a uint32) uint32 { return 0 })
	tf0(17, func(a uint32) uint32 { return 0 })

	tf(18, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return uint32(int32(a) % int32(b))
	})
	tf(19, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return a % b
	})
	tf0(18, func(a uint32) uint32 { return 0 })
	tf0(19, func(a uint32) uint32 { return 0 })

	tff(0, func(a, b float32) float32 { return a + b })
	tff(1, func(a, b float32) float32 { return a - b })
	tff(2, func(a, b float32) float32 { return a * b })
	tff(3, func(a, b float32) float32 { return a / b })
}
