package arch8

// IntBus is an interrupt bus for device
type IntBus interface {
	Ncore() byte
	Interrupt(code byte, core byte)
}

// IntAllCores generates a interrupt to all cores on the bus.
func IntAllCores(bus IntBus, code byte) {
	ncore := bus.Ncore()
	for i := byte(0); i < ncore; i++ {
		bus.Interrupt(code, i)
	}
}
