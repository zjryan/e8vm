package arch8

import (
	"testing"
)

func TestPhyMemory(t *testing.T) {
	eo := func(cond bool, s string, args ...interface{}) {
		if cond {
			t.Fatalf(s, args...)
		}
	}

	as := func(cond bool) {
		if !cond {
			t.Fatal()
		}
	}

	size := uint32(20 * PageSize)
	m := NewPhyMemory(size)
	eo(m.Size() != size, "size mismatch")

	w, e := m.ReadWord(4)
	eo(e != nil, "get an error for word reading")
	eo(w != 0, "page is not zeroed out")

	_, e = m.ReadWord(13)
	eo(e != errMisalign, "should have misalign error")
	_, e = m.ReadWord(size - 1)
	eo(e != errMisalign, "should have misalign error")

	_, e = m.ReadWord(size)
	eo(e != errOutOfRange, "should have out of range error")

	off := uint32(56 + PageSize*2)
	as(m.WriteByte(off+0, 0x37) == nil)
	as(m.WriteByte(off+1, 0x21) == nil)
	as(m.WriteByte(off+2, 0x5a) == nil)
	as(m.WriteByte(off+3, 0x70) == nil)
	exp := uint32(0x705a2137)
	w, e = m.ReadWord(off)
	as(e == nil)
	eo(w != exp, "expect 0x%08x got 0x%08x", exp, w)
	b, e := m.ReadByte(off + 2)
	as(e == nil)
	eo(b != 0x5a, "expect 0x5a got %02x", b)
}
