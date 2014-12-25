package arch8

// InstImm executes immediate instructions
type InstImm struct{}

// I executes the instruction.
// Might return invalid instruction exception,
// or memory related exceptions.
func (i *InstImm) I(cpu *CPU, in uint32) *Excep {
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
	case 1: // addi
		d = s + ims
	case 2: // slti
		if int32(s) < int32(ims) {
			d = 1
		} else {
			d = 0
		}
	case 3: // andi
		d = s & im
	case 4: // ori
		d = s | im
	case 5: // lui
		d = im << 16
	case 6: // lw
		d, e = cpu.virtMem.ReadWord(addr)
	case 7: // lb
		b, e = cpu.virtMem.ReadByte(addr)
		d = uint32(int32(int8(b)))
	case 8: // lbu
		b, e = cpu.virtMem.ReadByte(addr)
		d = uint32(b)
	case 9: // sw
		e = cpu.virtMem.WriteWord(addr, d)
	case 10: // sb
		e = cpu.virtMem.WriteByte(addr, byte(d))
	default:
		return errInvalidInst
	}

	if e != nil {
		return e
	}

	cpu.regs[dest] = d
	return nil
}
