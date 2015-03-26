package arch8

// InstSys exectues a system instruction
type instSys struct{}

// I executes the system instruction.
// Returns any exception encountered.
func (i *instSys) I(cpu *cpu, in uint32) *Excep {
	op := (in >> 24) & 0xff // (32:24]
	src := (in >> 21) & 0x7 // (24:21]
	s := cpu.regs[src]

	switch op {
	case HALT:
		return errHalt
	case SYSCALL:
		if !cpu.UserMode() {
			return errInvalidInst
		}
		return cpu.Syscall()
	case USERMOD:
		cpu.ring = 1
	case VTABLE:
		if cpu.UserMode() {
			return errInvalidInst
		}
		cpu.virtMem.SetTable(s)
	case IRET:
		if cpu.UserMode() {
			return errInvalidInst
		}
		return cpu.Iret()
	case CPUID:
		s = uint32(cpu.index)
	default:
		return errInvalidInst
	}

	cpu.regs[src] = s

	return nil
}
