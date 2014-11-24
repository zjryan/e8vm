package arch8

// IntBus is an interrupt bus for device
type IntBus interface {
	Ncore() int
	Interrupt(code byte, core byte)
}

// Device is a general interface of an pherical device.
type Device interface {
	Tick()
}
