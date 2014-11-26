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

	size := uint32(20 * PageSize)
	m := NewPhyMemory(size)
	eo(m.Size() != size, "size mismatch")
}
