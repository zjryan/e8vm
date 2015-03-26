package arch8

// InstJmp executes jump instruction
type instJmp struct{}

// I executes the jump instruction.
// Might return invalid instruction exception.
func (i *instJmp) I(cpu *cpu, in uint32) *Excep {
	op := (in >> 30) & 0x3 // (32:30]
	off := in & 0x3fffffff

	pc := cpu.regs[PC]
	to := pc + (off << 2)

	switch op {
	case J:
		/* do nothing */
	case JAL:
		cpu.regs[RET] = pc
	default:
		return errInvalidInst
	}

	cpu.regs[PC] = to
	return nil
}
