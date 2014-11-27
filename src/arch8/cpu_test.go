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
	as := func(cond bool, s string, args ...interface{}) {
		if !cond {
			t.Fatalf(s, args...)
		}
	}

	m := NewPhyMemory(PageSize * 32)
	cpu := NewCPU(m, ti1(0), 0)
	e := cpu.Tick()
	as(e == errHalt, "not halting")

	cpu.Reset()
	m.WriteWord(InitPC, 1)
	e = cpu.Tick()
	as(e == errTimeInt, "not time interrupt")

	cpu.Reset()
	cpu.interrupt.Enable()
	cpu.interrupt.EnableInt(errTimeInt.Code)
	cpu.interrupt.writeWord(intKernelSP, 0x10000)  // page 16
	cpu.interrupt.writeWord(intHandlerPC, 0x10000) // page 16

	e = cpu.Tick()
	as(e == nil, "should have no error")
	as(cpu.regs[PC] == 0x10000, "pc incorrect")
	as(cpu.regs[SP] == 0x10000, "sp incorrect")
	as(cpu.regs[RET] == 0x8000, "ret incorrect")
	b, e := m.ReadByte(0x10000 - intFrameSize + intFrameCode)
	as(e == nil, "read byte error")
	as(b == errTimeInt.Code, "not time interrupt")

	e = cpu.Tick()
	as(e == errHalt, "should halt")
}
