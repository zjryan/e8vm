package arch8

import (
	"bytes"
	"io"
)

// Machine is a multicore shared memory simulated arch8 machine.
type Machine struct {
	phyMem *PhyMemory
	inst   Inst
	cores  *MultiCore

	devices []Device
	serial  *Serial
	ticker  *Ticker
}

// NewMachine creates a machine with memory and cores.
func NewMachine(memSize uint32, ncore int) *Machine {
	ret := new(Machine)
	ret.phyMem = NewPhyMemory(memSize)
	ret.inst = new(InstArch8)
	ret.cores = NewMultiCore(ncore, ret.phyMem, ret.inst)

	// hook-up devices
	p := ret.phyMem.P(pageBasicIO)
	ret.serial = NewSerial(p, ret.cores)
	ret.ticker = NewTicker(ret.cores)
	ret.AddDevice(ret.serial)
	ret.AddDevice(ret.ticker)

	return ret
}

// AddDevice adds a devices to the machine.
func (m *Machine) AddDevice(d Device) {
	m.devices = append(m.devices, d)
}

// Tick proceeds the simulation by one tick.
func (m *Machine) Tick() *Excep {
	for _, d := range m.devices {
		d.Tick()
	}

	return m.cores.Tick()
}

// Run simulates nticks. It returns the number of ticks
// simulated without error, and the first met error if any.
func (m *Machine) Run(nticks int) (int, *Excep) {
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
		p := m.phyMem.P(pn)
		if p == nil {
			return errOutOfRange
		}
		buf := bytes.NewBuffer(p.Bytes)
		_, e := io.CopyN(buf, r, PageSize)
		if e == io.EOF {
			return nil
		} else if e != nil {
			return e
		}
		pn++
	}
}
