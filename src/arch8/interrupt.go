package arch8

// Interrupt defines the interrupt page
type Interrupt struct {
	p *Page // the dma page for interrupt handler
}

const Ninterrupt = 256

const (
	intBase      = 0            // base offset in the page
	intFlags     = intBase + 0  // flags, bit 0 is master enabling switch
	intKernelSP  = intBase + 4  // position of the handler stack base pointer
	intHandlerPC = intBase + 8  // position of the handler start PC
	intMask      = intBase + 32 // interrupt enable mask bits offset (32 bytes)
	intPending   = intBase + 64 // interrupt pending bits offset (32 bytes)
)

// NewInterrupt creates a interrupt on the given DMA page.
func NewInterrupt(p *Page) *Interrupt {
	ret := new(Interrupt)
	ret.p = p

	return ret
}

// Issue issues an interrupt. If the interrupt is already issued,
// this has no effect.
func (in *Interrupt) Issue(i byte) {
	off := uint32(i/8) + intMask
	b := in.p.ReadByte(off)
	b |= 0x1 << (i % 8)
	in.p.WriteByte(off, b)
}

// Clear clears an interrupt. If the interrupt is already cleared,
// this has no effect.
func (in *Interrupt) Clear(i byte) {
	off := uint32(i/8) + intMask
	b := in.p.ReadByte(off)
	b &= ^(0x1 << (i % 8))
	in.p.WriteByte(off, b)
}

// Enable sets the interrupt enable bit in the flags.
func (in *Interrupt) Enable() {
	b := in.p.ReadByte(intFlags)
	b |= 0x1
	in.p.WriteByte(intFlags, b)
}

// Disable clears the interrupt enable bit in the flags.
func (in *Interrupt) Disable() {
	b := in.p.ReadByte(intFlags)
	b &= ^byte(0x1)
	in.p.WriteByte(intFlags, b)
}

// EnableInt enables a particular interrupt.
func (in *Interrupt) EnableInt(i byte) {
	off := uint32(i/8) + intPending
	b := in.p.ReadByte(off)
	b |= 0x1 << (i % 8)
	in.p.WriteByte(off, b)
}

// DisableInt disables a particular interrupt.
func (in *Interrupt) DisableInt(i byte) {
	off := uint32(i/8) + intMask
	b := in.p.ReadByte(off)
	b &= ^(0x1 << (i % 8))
	in.p.WriteByte(off, b)
}

// Flags returns the flags byte.
func (in *Interrupt) Flags() byte {
	return in.p.ReadByte(intFlags)
}

// Poll looks for the next pending interrupt.
func (in *Interrupt) Poll() (bool, byte) {
	flag := in.Flags()
	if flag&0x1 == 0 { // interrupt is disabled
		return false, 0
	}

	// search bits based on priorities.
	// smaller is higher
	for i := uint32(0); i < Ninterrupt/8; i++ {
		pending := in.p.ReadByte(intPending + i)
		mask := in.p.ReadByte(intMask + i)
		pending &= mask
		if pending == 0 {
			continue
		}

		for b := byte(0); b < 8; b++ {
			if pending&(0x1<<b) == 0 {
				continue
			}

			return true, byte(i*8) + b
		}
	}

	return false, 0
}
