package arch8

import (
	"testing"

	"math"
	"math/rand"
)

func TestInstReg(t *testing.T) {
	m := newPhyMemory(PageSize * 32)
	cpu := newCPU(m, new(instReg), 0)

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

	tf(SLLV, func(a, b uint32) uint32 { return a << b })
	tf(SRLV, func(a, b uint32) uint32 { return a >> b })
	tf(SRLA, func(a, b uint32) uint32 { return uint32(int32(a) >> b) })

	tfs(SLL, func(a, sh uint32) uint32 { return a << sh })
	tfs(SRL, func(a, sh uint32) uint32 { return a >> sh })
	tfs(SRA, func(a, sh uint32) uint32 { return uint32(int32(a) >> sh) })

	tf(ADD, func(a, b uint32) uint32 { return a + b })
	tf(SUB, func(a, b uint32) uint32 { return a - b })
	tf(AND, func(a, b uint32) uint32 { return a & b })
	tf(OR, func(a, b uint32) uint32 { return a | b })
	tf(XOR, func(a, b uint32) uint32 { return a ^ b })
	tf(NOR, func(a, b uint32) uint32 { return ^(a | b) })
	tf(SLT, func(a, b uint32) uint32 {
		if int32(a) < int32(b) {
			return 1
		}
		return 0
	})
	tf(SLTU, func(a, b uint32) uint32 {
		if a < b {
			return 1
		}
		return 0
	})
	tf(MUL, func(a, b uint32) uint32 {
		return uint32(int32(a) * int32(b))
	})
	tf(MULU, func(a, b uint32) uint32 { return a * b })

	tf(DIV, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return uint32(int32(a) / int32(b))
	})
	tf(DIVU, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return a / b
	})
	tf0(DIV, func(a uint32) uint32 { return 0 })
	tf0(DIVU, func(a uint32) uint32 { return 0 })

	tf(MOD, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return uint32(int32(a) % int32(b))
	})
	tf(MODU, func(a, b uint32) uint32 {
		if b == 0 {
			return 0
		}
		return a % b
	})
	tf0(MOD, func(a uint32) uint32 { return 0 })
	tf0(MODU, func(a uint32) uint32 { return 0 })

	tff(FADD, func(a, b float32) float32 { return a + b })
	tff(FSUB, func(a, b float32) float32 { return a - b })
	tff(FMUL, func(a, b float32) float32 { return a * b })
	tff(FDIV, func(a, b float32) float32 { return a / b })
}
