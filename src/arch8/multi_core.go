package arch8

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

// Tick performs one tick on each core.
func (c *MultiCore) Tick() *Excep {
	for i, core := range c.cores {
		e := core.Tick()
		if e != nil {
			e.Core = byte(i)
			return e
		}
	}

	return nil
}
