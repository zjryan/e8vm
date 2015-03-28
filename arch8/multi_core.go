package arch8

import (
	"fmt"
)

// MultiCore simulates a shared memory multicore processor.
type multiCore struct {
	cores  []*cpu
	phyMem *phyMemory
}

// NewMultiCore creates a shared memory multicore processor.
func newMultiCore(ncore int, mem *phyMemory, i inst) *multiCore {
	if ncore > 32 {
		panic("too many cores")
	}

	ret := new(multiCore)
	ret.cores = make([]*cpu, ncore)
	ret.phyMem = mem

	for ind := range ret.cores {
		ret.cores[ind] = newCPU(mem, i, byte(ind))
	}

	return ret
}

// CoreExcep is an exception on a particular core.
type CoreExcep struct {
	Core int
	*Excep
}

// Tick performs one tick on each core.
func (c *multiCore) Tick() *CoreExcep {
	for i, core := range c.cores {
		e := core.Tick()
		if e != nil {
			return &CoreExcep{i, e}
		}
	}

	return nil
}

// Ncore returns the number of cores.
func (c *multiCore) Ncore() byte {
	return byte(len(c.cores))
}

// Interrupt issues an interrupt to a particular core.
func (c *multiCore) Interrupt(code byte, core byte) {
	if int(core) >= len(c.cores) {
		panic("out of cores")
	}

	c.cores[core].Interrupt(code)
}

// PrintStatus prints out the core status of all the cores.
func (c *multiCore) PrintStatus() {
	for i, core := range c.cores {
		if len(c.cores) > 1 {
			fmt.Printf("[core %d]\n", i)
		}
		printCPUStatus(core)
		fmt.Println()
	}
}

func printCPUStatus(c *cpu) {
	p := func(name string, reg int) {
		fmt.Printf(" %3s = 0x%08x %-11d\n",
			name, c.regs[reg], int32(c.regs[reg]),
		)
	}

	p("r0", R0)
	p("r1", R1)
	p("r2", R2)
	p("r3", R3)
	p("r4", R4)
	p("sp", SP)
	p("ret", RET)
	p("pc", PC)

	fmt.Printf("ring = %d\n", c.ring)
}
