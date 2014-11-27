package arch8

import (
	"testing"
)

func TestInstReg(t *testing.T) {
	m := NewPhyMemory(PageSize * 32)
	cpu := NewCPU(m, new(InstReg), 0)

	cpu.Tick()
}
