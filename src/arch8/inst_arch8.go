package arch8

// InstArch8 dispatches and executes an arch8 instruction.
type InstArch8 struct {
	reg InstReg
	imm InstImm
	br  InstBr
	sys InstSys
	jmp InstJmp
}

// I executes an arch8 instructino
func (i *InstArch8) I(cpu *CPU, in uint32) *Excep {
	if (in >> 31) == 0 {
		op := (in >> 24) & 0xff
		switch {
		case op == 0: // op == 0
			return i.reg.I(cpu, in)
		case op < 32: // op in (0, 32)
			return i.imm.I(cpu, in)
		case op < 64: // op in [32, 64)
			return i.br.I(cpu, in)
		case op < 128: // op in [64, 128)
			return i.sys.I(cpu, in)
		default:
			panic("bug")
		}
	}

	return i.jmp.I(cpu, in)
}
