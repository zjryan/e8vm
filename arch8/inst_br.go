package arch8

// InstBr exectues branch instruction
type instBr struct{}

// I executes the instruction.
// Might return invalid instruction exception.
func (i *instBr) I(cpu *cpu, in uint32) *Excep {
	op := (in >> 24) & 0xff  // (32:24]
	src1 := (in >> 21) & 0x7 // (24:21]
	src2 := (in >> 18) & 0x7 // (21:18]
	im := in & 0x3ffff       // (18:0]

	s1 := cpu.regs[src1]
	s2 := cpu.regs[src2]
	pc := cpu.regs[PC]
	br := pc + uint32(int32(im<<14)>>12)

	switch op {
	case BNE:
		if s1 != s2 {
			pc = br
		}
	case BEQ:
		if s1 == s2 {
			pc = br
		}
	default:
		return errInvalidInst
	}

	cpu.regs[PC] = pc
	return nil
}
