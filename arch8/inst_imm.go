package arch8

// InstImm executes immediate instructions
type instImm struct{}

// I executes the instruction.
// Might return invalid instruction exception,
// or memory related exceptions.
func (i *instImm) I(cpu *cpu, in uint32) *Excep {
	op := (in >> 24) & 0xff  // (32:24]
	dest := (in >> 21) & 0x7 // (24:21]
	src := (in >> 18) & 0x7  // (21:18]
	im := in & 0xffff        // (16:0]

	s := cpu.regs[src]
	d := cpu.regs[dest]
	ims := uint32(int32(im<<16) >> 16)
	addr := s + ims
	var e *Excep
	var b byte

	switch op {
	case 0:
		panic("register based instruction")
	case ADDI:
		d = s + ims
	case SLTI:
		if int32(s) < int32(ims) {
			d = 1
		} else {
			d = 0
		}
	case ANDI:
		d = s & im
	case ORI:
		d = s | im
	case XORI:
		d = s ^ im
	case LUI:
		d = im << 16
	case LW:
		d, e = cpu.readWord(addr)
	case LB:
		b, e = cpu.readByte(addr)
		d = uint32(int32(int8(b)))
	case LBU:
		b, e = cpu.readByte(addr)
		d = uint32(b)
	case SW:
		e = cpu.writeWord(addr, d)
	case SB:
		e = cpu.writeByte(addr, byte(d))
	default:
		return errInvalidInst
	}

	if e != nil {
		return e
	}

	cpu.regs[dest] = d
	return nil
}
