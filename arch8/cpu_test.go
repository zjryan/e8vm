package arch8

import (
	"testing"

	"lonnie.io/e8vm/conv"
)

type ti1 int

func (i ti1) I(cpu *cpu, in uint32) *Excep {
	switch in {
	case 0:
		return errHalt
	case 1:
		return errTimeInt
	case 2: // iret
		if cpu.UserMode() {
			return errInvalidInst
		}
		return cpu.Iret()
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

	m := newPhyMemory(PageSize * 32)
	cpu := newCPU(m, ti1(0), 0)
	e := cpu.Tick()
	as(e == errHalt, "not halting")

	cpu.Reset()
	m.WriteWord(conv.InitPC, 1)
	e = cpu.Tick()
	as(e == errTimeInt, "not time interrupt")

	cpu.Reset()
	cpu.interrupt.Enable()
	cpu.interrupt.EnableInt(errTimeInt.Code)
	cpu.interrupt.writeWord(intKernelSP, 0x10000)  // page 16
	cpu.interrupt.writeWord(intHandlerPC, 0x10000) // page 16
	cpu.ring = 1

	e = cpu.Tick()
	as(e == nil, "should have no error")
	as(cpu.regs[PC] == 0x10000, "pc incorrect")
	as(cpu.regs[SP] == 0x10000, "sp incorrect")
	as(cpu.regs[RET] == 0x8000, "ret incorrect")
	as(!cpu.UserMode(), "not in kernel")
	b, e := m.ReadByte(0x10000 - intFrameSize + intFrameCode)
	as(e == nil, "read byte error")
	as(b == errTimeInt.Code, "not time interrupt")
	b, e = m.ReadByte(0x10000 - intFrameSize + intFrameRing)
	as(b == 1, "was not in user")
	as(!cpu.interrupt.Enabled(), "interrupt not disabled")

	e = cpu.Tick()
	as(e == errHalt, "should halt now")

	cpu.Reset()
	cpu.interrupt.Enable()
	cpu.ring = 1
	m.WriteWord(0x10000, 2) // write an iret
	e = cpu.Tick()
	as(e == nil, "unexpected error: %s", e)
	as(!cpu.interrupt.Enabled(), "interrupt not disabled")
	e = cpu.Tick()
	as(e == nil, "unexpected error: %s", e)
	as(cpu.regs[PC] == 0x8000, "pc not iret'ed")
	as(cpu.regs[SP] == 0, "sp not restored")
	as(cpu.ring == 1, "ring not restored")
	as(cpu.interrupt.Enabled(), "interrupt not enabled again")
	has, code := cpu.interrupt.Poll()
	as(!has && code == 0, "interrupt was not cleared")
}
