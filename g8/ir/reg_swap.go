package ir

func saveVar(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.sw(reg, _sp, v.offset))
	} else if v.size == 1 {
		b.inst(asm.sb(reg, _sp, v.offset))
	} else {
		panic("invalid size to save from a register")
	}
}

func loadVar(b *Block, reg uint32, v *stackVar) {
	if v.size == regSize {
		b.inst(asm.lw(reg, _sp, v.offset))
	} else if v.size == 1 {
		b.inst(asm.lb(reg, _sp, v.offset))
	} else {
		panic("invalid size to load to a register")
	}
}

func saveRef(b *Block, reg uint32, r ref) {
	switch r := r.(type) {
	case *stackVar:
		saveVar(b, reg, r)
	case *number:
		panic("numbers are read only")
	default:
		panic("not implemented")
	}
}

func loadRef(b *Block, reg uint32, r ref) {
	switch r := r.(type) {
	case *stackVar:
		loadVar(b, reg, r)
	case *number:
		high := r.v >> 16
		if high != 0 {
			b.inst(asm.lui(reg, high))
		}
		b.inst(asm.ori(reg, reg, r.v))
	default:
		panic("not implemented")
	}
}
