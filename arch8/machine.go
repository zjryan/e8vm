package arch8

import (
	"io"
)

// Machine is a multicore shared memory simulated arch8 machine.
type Machine struct {
	phyMem *phyMemory
	inst   inst
	cores  *multiCore

	devices []device
}

// NewMachine creates a machine with memory and cores.
func NewMachine(memSize uint32, ncore int) *Machine {
	ret := new(Machine)
	ret.phyMem = newPhyMemory(memSize)
	ret.inst = new(instArch8)
	ret.cores = newMultiCore(ncore, ret.phyMem, ret.inst)

	// hook-up devices
	p := ret.phyMem.Page(pageBasicIO)

	ret.addDevice(newTicker(ret.cores))
	ret.addDevice(newSerial(p, ret.cores))
	ret.addDevice(newConsole(p, ret.cores))

	return ret
}

// AddDevice adds a devices to the machine.
func (m *Machine) addDevice(d device) {
	m.devices = append(m.devices, d)
}

// Tick proceeds the simulation by one tick.
func (m *Machine) Tick() *CoreExcep {
	for _, d := range m.devices {
		d.Tick()
	}

	return m.cores.Tick()
}

// Run simulates nticks. It returns the number of ticks
// simulated without error, and the first met error if any.
func (m *Machine) Run(nticks int) (int, *CoreExcep) {
	n := 0
	for i := 0; nticks == 0 || i < nticks; i++ {
		e := m.Tick()
		n++
		if e != nil {
			return n, e
		}
	}

	return n, nil
}

// LoadImage loads an booting image from a reader and put it
// at a particular offset in the physical memory.
func (m *Machine) LoadImage(r io.Reader, offset uint32) error {
	if offset%PageSize != 0 {
		panic("boot image not page aligned")
	}

	pn := offset / PageSize

	for {
		p := m.phyMem.Page(pn)
		if p == nil {
			return errOutOfRange
		}
		_, e := r.Read(p.Bytes)
		if e == io.EOF {
			return nil
		} else if e != nil {
			return e
		}

		pn++
	}
}

// PrintCoreStatus prints the cpu statuses.
func (m *Machine) PrintCoreStatus() {
	m.cores.PrintStatus()
}
