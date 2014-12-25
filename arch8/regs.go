package arch8

// Nreg is the number of registers
const Nreg = 8

// Register names
const (
	R0 = 0
	R1 = 1
	R2 = 2
	R3 = 3
	R4 = 4

	SP  = 5
	RET = 6
	PC  = 7
)

func makeRegs() []uint32 {
	return make([]uint32, Nreg)
}
