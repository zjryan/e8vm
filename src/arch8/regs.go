package arch8

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

	Nreg = 8 // total number of registers
)

func makeRegs() []uint32 {
	return make([]uint32, Nreg)
}
