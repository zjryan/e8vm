package arch8

import (
	"testing"
)

type ti1 int

func (i ti1) I(cpu *CPU, in uint32) *Excep {
	switch in {
	case 0:
		return errHalt
	case 1:
		return errTimeInt
	default:
		return errInvalidInst
	}
}

func TestCPU(t *testing.T) {
	as := func(cond bool) {
		if !cond {
			t.Fail()
		}
	}

	m := NewPhyMemory(PageSize * 32)
	cpu := NewCPU(m, ti1(0), 0)
	e := cpu.Tick()
	as(e == errHalt)
}
