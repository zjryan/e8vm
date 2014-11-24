package arch8

// Inst is an interface for executing one single instruction
type Inst interface {
	I(cpu *CPU, in uint32) *Excep
}

// CPU defines the structure of a processing unit.
type CPU struct {
	regs []uint32

	ring      byte
	phyMem    *PhyMemory
	virtMem   *VirtMemory
	interrupt *Interrupt

	inst Inst

	clock uint32 // cycle counter
}

// NewCPU creates a CPU that does not have a instruction binding
func NewCPU(memSize uint32) *CPU {
	ret := new(CPU)
	ret.regs = makeRegs()
	ret.phyMem = NewPhyMemory(memSize)
	ret.virtMem = NewVirtMemory(ret.phyMem)

	intPage := ret.phyMem.P(1)
	if intPage == nil {
		panic("memory too small")
	}
	ret.interrupt = NewInterrupt(intPage)

	return ret
}

func (c *CPU) tick() *Excep {
	pc := c.regs[PC]
	inst, e := c.virtMem.ReadWord(pc)
	if e != nil {
		return e
	}

	c.regs[PC] = pc + 4
	if c.inst != nil {
		e = c.inst.I(c, inst)
		if e != nil {
			c.regs[PC] = pc // restore saved PC
			return e
		}
	}

	return nil
}

var regsToPush = []int{SP, PC}

const (
	intFrameSize = 12

	intFrameSP   = 0
	intFrameRET  = 4
	intFrameCode = 8
	intFrameRing = 9
)

// Interrupt sets up a interrupt routine.
func (c *CPU) Interrupt(code byte) *Excep {
	ksp := c.interrupt.kernelSP()
	base := ksp - intFrameSize

	if e := c.virtMem.WriteWord(base+intFrameSP, c.regs[SP]); e != nil {
		return e
	}
	if e := c.virtMem.WriteWord(base+intFrameRET, c.regs[RET]); e != nil {
		return e
	}
	if e := c.virtMem.WriteByte(base+intFrameCode, code); e != nil {
		return e
	}
	if e := c.virtMem.WriteByte(base+intFrameRing, c.ring); e != nil {
		return e
	}

	c.interrupt.Disable()
	c.regs[SP] = ksp
	c.regs[RET] = c.regs[PC]
	c.regs[PC] = c.interrupt.handlerPC()
	c.ring = 0

	return nil
}

// Iret restores from an interrupt
func (c *CPU) Iret() *Excep {
	ksp := c.interrupt.kernelSP()
	base := ksp - intFrameSize
	sp, e := c.virtMem.ReadWord(base + intFrameSP)
	if e != nil {
		return e
	}
	ret, e := c.virtMem.ReadWord(base + intFrameRET)
	if e != nil {
		return e
	}
	code, e := c.virtMem.ReadByte(base + intFrameCode)
	if e != nil {
		return e
	}
	ring, e := c.virtMem.ReadByte(base + intFrameRing)
	if e != nil {
		return e
	}

	c.regs[PC] = c.regs[RET]
	c.regs[RET] = ret
	c.regs[SP] = sp
	c.ring = ring
	c.interrupt.Clear(code)
	c.interrupt.Enable()

	return nil
}

// Tick executes one instruction, and increases the program counter
// by 4 by default. If an exception is met, it will handle it.
func (c *CPU) Tick() *Excep {
	poll, code := c.interrupt.Poll()
	if poll {
		return c.Interrupt(code)
	}

	// no interrupt to dispatch, let's proceed
	e := c.Tick()
	if e != nil {
		// proceed attempt failed, handle the error
		c.interrupt.Issue(e.Code)
		poll, code := c.interrupt.Poll()
		if poll {
			if code != e.Code {
				panic("interrupt code is different")
			}
			return c.Interrupt(code)
		}
	}

	return e
}
