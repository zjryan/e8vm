package link8

// link is a piece of linking information for an instruction
// in a code section at a particular offset.
// it uses the indices in the package for symbol lookup
type link struct {
	offset uint32
	pkg    uint32 // relative package index
	sym    uint32
}

// Func is a relocatable code section
type Func struct {
	insts []uint32
	links []*link

	addr uint32
}

// NewFunc creates a new relocatable code section.
func NewFunc() *Func {
	return new(Func)
}

// AddInst appends an instruction at the end of the function.
func (f *Func) AddInst(i uint32) {
	f.insts = append(f.insts, i)
}

// TooLarge checks if the function size is larger than 2GB.
func (f *Func) TooLarge() bool {
	return len(f.insts)*4 >= 0x80000000
}

// Size returns the size of the function.
func (f *Func) Size() uint32 {
	return uint32(len(f.insts) * 4)
}

// AddLink links the last instruction in inst to
// the symbol pkg.sym, where pkg and sym are using the
// indexing of the object file.
// fill field must be less than 4 so that it fits in the
// lowest 2 bits in the offset field. The other bits
// of the offset fields will be automatically calculated
// based on the number of instructions in insts.
func (f *Func) addLink(fill int, pkg, sym uint32) {
	if len(f.insts) == 0 {
		panic("no inst to link")
	}
	if !(fill > 0 && fill <= 3) {
		panic("invalid fill")
	}

	offset := uint32(len(f.insts))*4 - 4
	offset |= uint32(fill) & 0x3
	link := &link{offset: offset, pkg: pkg, sym: sym}
	f.links = append(f.links, link)
}

const (
	FillNone = iota
	FillLink
	FillLow
	FillHigh
)
