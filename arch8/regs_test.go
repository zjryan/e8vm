package arch8

import (
	"testing"
)

func TestRegs(t *testing.T) {
	regs := makeRegs()
	if len(regs) != Nreg {
		t.Error("unmatched number of regs")
	}
}
