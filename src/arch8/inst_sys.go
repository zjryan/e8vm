package arch8

// InstSys exectues a system instruction
type InstSys struct{}

// I executes the system instruction.
// Returns any exception encountered.
func (i *InstSys) I(cpu *CPU, in uint32) *Excep {
	op := (in >> 24) & 0xff // (32:24]
	src := (in >> 21) & 0x7 // (24:21]
	s := cpu.regs[src]

	switch op {
	case 64: // halt
		return errHalt
	case 65: // syscall
		if !cpu.UserMode() {
			return errInvalidInst
		}
		return cpu.Syscall()
	case 66: // usermod
		cpu.virtMem.Ring = 1
	case 67: // vtable
		if cpu.UserMode() {
			return errInvalidInst
		}
		cpu.virtMem.SetTable(s)
	case 68: // iret
		if cpu.UserMode() {
			return errInvalidInst
		}
		return cpu.Iret()
	case 69: // cpuid
		s = uint32(cpu.index)
	default:
		return errInvalidInst
	}

	cpu.regs[src] = s

	return nil
}
