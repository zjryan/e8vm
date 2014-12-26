package arch8

import (
	"fmt"
)

// MultiCore simulates a shared memory multicore processor.
type MultiCore struct {
	cores  []*CPU
	phyMem *PhyMemory
}

// NewMultiCore creates a shared memory multicore processor.
func NewMultiCore(ncore int, mem *PhyMemory, i Inst) *MultiCore {
	if ncore > 32 {
		panic("too many cores")
	}

	ret := new(MultiCore)
	ret.cores = make([]*CPU, ncore)
	ret.phyMem = mem

	for ind := range ret.cores {
		ret.cores[ind] = NewCPU(mem, i, byte(ind))
	}

	return ret
}

// CoreExcep is an exception on a particular core.
type CoreExcep struct {
	Core int
	*Excep
}

// Tick performs one tick on each core.
func (c *MultiCore) Tick() *CoreExcep {
	for i, core := range c.cores {
		e := core.Tick()
		if e != nil {
			return &CoreExcep{i, e}
		}
	}

	return nil
}

// Ncore returns the number of cores.
func (c *MultiCore) Ncore() byte {
	return byte(len(c.cores))
}

// Interrupt issues an interrupt to a particular core.
func (c *MultiCore) Interrupt(code byte, core byte) {
	if int(core) >= len(c.cores) {
		panic("out of cores")
	}

	c.cores[core].Interrupt(code)
}

// PrintStatus prints out the core status of all the cores.
func (c *MultiCore) PrintStatus() {
	for i, core := range c.cores {
		if len(c.cores) > 1 {
			fmt.Printf("[core %d]\n", i)
		}
		printCPUStatus(core)
		fmt.Println()
	}
}

func printCPUStatus(c *CPU) {
	p := func(name string, reg int) {
		fmt.Printf(" %3s = 0x%08x %-11d\n", name, c.regs[reg], int32(c.regs[reg]))
	}

	p("r0", R0)
	p("r1", R1)
	p("r2", R2)
	p("r3", R3)
	p("r4", R4)
	p("sp", SP)
	p("ret", RET)
	p("pc", PC)

	fmt.Printf("ring = %d\n", c.virtMem.Ring)
}
