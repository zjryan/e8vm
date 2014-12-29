package asm8

type link struct {
	offset uint32
	pkg    uint32 // relative package index
	sym    uint32
}

type funcObj struct {
	insts []uint32
	links []*link
}

func (o *funcObj) addInst(i uint32) {
	o.insts = append(o.insts, i)
}

// addLink links the last instruction in inst to
// the symbol pkg.sym, where pkg and sym are using the
// indexing of the object file.
// fill field must be less than 4 so that it fits in the
// lowest 2 bits in the offset field. The other bits
// of the offset fields will be automatically calculated
// based on the number of instructions in insts.
func (o *funcObj) addLink(fill int, pkg, sym uint32) {
	if len(o.insts) == 0 {
		panic("no inst to link")
	}
	if !(fill > 0 && fill <= 3) {
		panic("invalid fill")
	}

	offset := uint32(len(o.insts))*4 - 4
	offset |= uint32(fill) & 0x3
	link := &link{offset: offset, pkg: pkg, sym: sym}
	o.links = append(o.links, link)
}
