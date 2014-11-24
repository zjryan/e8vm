package arch8

// CPU defines the structure of a processing unit.
type CPU struct {
	regs      []uint32
	phyMem    *PhyMemory
	virtMem   *VirtMemory
	interrupt *Interrupt

	inst func(cpu *CPU, in uint32) *Excep // instruction executor

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
		e = c.inst(c, inst)
		if e != nil {
			c.regs[PC] = pc // restore saved PC
			return e
		}
	}

	return nil
}

// Tick executes one instruction, and increases the program counter
// by 4 by default. If an exception is met, it will handle it.
func (c *CPU) Tick() *Excep {
	e := c.Tick()
	if e != nil {
		panic("todo")
	}

	return nil
}
