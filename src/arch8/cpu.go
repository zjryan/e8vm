package arch8

type CPU struct {
	regs      []uint32
	phyMem    *PhyMemory
	virtMem   *VirtMemory
	interrupt *Interrupt

	inst func(cpu *CPU, in uint32) error // instruction executor

	clock uint32 // cycle counter
}

// Creates a CPU that does not have a instruction binding
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

// Tick executes one instruction, and increases the program counter
// by 4 by default.
func (c *CPU) tick() error {
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
