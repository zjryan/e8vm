package arch8

// Inst is an interface for executing one single instruction
type Inst interface {
	I(cpu *CPU, in uint32) *Excep
}

// CPU defines the structure of a processing unit.
type CPU struct {
	regs []uint32
	ring byte

	phyMem    *PhyMemory
	virtMem   *VirtMemory
	interrupt *Interrupt

	inst  Inst
	index byte
}

// InitPC points the default starting value of the program counter.
const InitPC = 0x8000

// NewCPU creates a CPU with memroy and instruction binding
func NewCPU(mem *PhyMemory, i Inst, index byte) *CPU {
	if index >= 32 {
		panic("too many cores")
	}

	ret := new(CPU)
	ret.regs = makeRegs()
	ret.phyMem = mem
	ret.virtMem = NewVirtMemory(ret.phyMem)
	ret.index = index

	intPage := ret.phyMem.Page(pageInterrupt) // page 1 is the interrupt page
	if intPage == nil {
		panic("memory too small")
	}
	ret.interrupt = NewInterrupt(intPage, index)
	ret.inst = i

	ret.regs[PC] = InitPC

	return ret
}

// UserMode returns trun when the CPU is in user mode.
func (c *CPU) UserMode() bool {
	return c.ring > 0
}

// Reset resets the CPU's internal states, i.e., registers,
// the page table, and disables interrupt
func (c *CPU) Reset() {
	for i := 0; i < Nreg; i++ {
		c.regs[i] = 0
	}
	c.regs[PC] = InitPC
	c.virtMem.SetTable(0)
	c.ring = 0
	c.interrupt.Disable()
}

func (c *CPU) tick() *Excep {
	pc := c.regs[PC]
	inst, e := c.readWord(pc)
	if e != nil {
		return e
	}

	c.regs[PC] = pc + 4
	if c.inst != nil {
		e = c.inst.I(c, inst)

		if e != nil {
			c.regs[PC] = pc // restore saved original PC
			return e
		}
	}

	return nil
}

const (
	intFrameSP   = 0
	intFrameRET  = 4
	intFrameArg  = 8
	intFrameCode = 12
	intFrameRing = 13

	intFrameSize = 16
)

// Interrupt issues an interrupt to the core
func (c *CPU) Interrupt(code byte) {
	c.interrupt.Issue(code)
}

func (c *CPU) readWord(addr uint32) (uint32, *Excep) {
	return c.virtMem.ReadWord(addr, c.ring)
}

func (c *CPU) readByte(addr uint32) (uint8, *Excep) {
	return c.virtMem.ReadByte(addr, c.ring)
}

func (c *CPU) writeWord(addr uint32, v uint32) *Excep {
	return c.virtMem.WriteWord(addr, c.ring, v)
}

func (c *CPU) writeByte(addr uint32, v uint8) *Excep {
	return c.virtMem.WriteByte(addr, c.ring, v)
}

// Ienter enters a interrupt routine.
func (c *CPU) Ienter(code byte, arg uint32) *Excep {
	ksp := c.interrupt.kernelSP()
	base := ksp - intFrameSize

	writeWord := func(off uint32, v uint32) *Excep {
		return c.virtMem.WriteWord(base+off, 0, v)
	}
	writeByte := func(off uint32, b uint8) *Excep {
		return c.virtMem.WriteByte(base+off, 0, b)
	}
	if e := writeWord(intFrameSP, c.regs[SP]); e != nil {
		return e
	}
	if e := writeWord(intFrameRET, c.regs[RET]); e != nil {
		return e
	}
	if e := writeWord(intFrameArg, arg); e != nil {
		return e
	}
	if e := writeByte(intFrameCode, code); e != nil {
		return e
	}
	if e := writeByte(intFrameRing, c.ring); e != nil {
		return e
	}

	c.interrupt.Disable()
	c.regs[SP] = ksp
	c.regs[RET] = c.regs[PC]
	c.regs[PC] = c.interrupt.handlerPC()
	c.ring = 0

	return nil
}

// Syscall jumps to the system call handler and switches to ring 0.
func (c *CPU) Syscall() *Excep {
	c.regs[PC] = c.interrupt.syscallPC()
	c.ring = 0
	return nil
}

// Iret restores from an interrupt.
// It restores the SP, RET, PC registers, restores the ring level,
// clears the served interrupt bit and enables interrupt again.
func (c *CPU) Iret() *Excep {
	if c.ring != 0 {
		panic("iret in userland")
	}

	ksp := c.interrupt.kernelSP()
	base := ksp - intFrameSize
	sp, e := c.readWord(base + intFrameSP)
	if e != nil {
		return e
	}
	ret, e := c.readWord(base + intFrameRET)
	if e != nil {
		return e
	}
	code, e := c.readByte(base + intFrameCode)
	if e != nil {
		return e
	}
	ring, e := c.readByte(base + intFrameRing)
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
		return c.Ienter(code, 0)
	}

	// no interrupt to dispatch, let's proceed
	e := c.tick()
	if e != nil {
		// proceed attempt failed, handle the error
		c.interrupt.Issue(e.Code)
		poll, code := c.interrupt.Poll()
		if poll {
			if code != e.Code {
				panic("interrupt code is different")
			}
			return c.Ienter(code, e.Arg)
		}
	}

	return e
}
