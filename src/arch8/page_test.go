package arch8

import (
	"testing"
)

func TestPage(t *testing.T) {
	e := func(cond bool, s string, args ...interface{}) {
		if cond {
			t.Fatalf(s, args...)
		}
	}

	e(PageSize < 4, "page size too small")
	p := NewPage()
	e(len(p.Bytes) != PageSize, "page size incorrect")
	e(((PageSize-1)&PageSize) != 0, "page size not exp of 2")

	for i := uint32(0); i < PageSize; i++ {
		b := p.ReadByte(i)
		e(b != 0, "byte %d not zero on new page", i)
	}

	for i := uint32(0); i < PageSize/4; i++ {
		b := p.ReadWord(i * 4)
		e(b != 0, "word %d not zero on new page", i)
	}
}
