package arch8

// Interrupt defines the interrupt page
type Interrupt struct {
	p    *Page  // the dma page for interrupt handler
	base uint32 // dma offset
}

// Number of interrupts
const Ninterrupt = 256

const (
	intFlags     = 0  // flags, bit 0 is master enabling switch
	intKernelSP  = 4  // position of the handler stack base pointer
	intHandlerPC = 8  // position of the handler start PC
	intSyscallPC = 12 // position of the syscall start PC
	intMask      = 32 // interrupt enable mask bits offset (32 bytes)
	intPending   = 64 // interrupt pending bits offset (32 bytes)

	intCtrlSize = 128
)

// NewInterrupt creates a interrupt on the given DMA page.
func NewInterrupt(p *Page, core byte) *Interrupt {
	ret := new(Interrupt)
	ret.p = p
	ret.base = uint32(core) * intCtrlSize
	if ret.base+intCtrlSize > PageSize {
		panic("bug")
	}

	return ret
}

func (in *Interrupt) readWord(off uint32) uint32 {
	return in.p.ReadWord(in.base + off)
}

func (in *Interrupt) readByte(off uint32) byte {
	return in.p.ReadByte(in.base + off)
}

func (in *Interrupt) writeWord(off uint32, v uint32) {
	in.p.WriteWord(in.base+off, v)
}

func (in *Interrupt) writeByte(off uint32, v byte) {
	in.p.WriteByte(in.base+off, v)
}

func (in *Interrupt) kernelSP() uint32 {
	return in.readWord(intKernelSP)
}

func (in *Interrupt) handlerPC() uint32 {
	return in.readWord(intHandlerPC)
}

func (in *Interrupt) syscallPC() uint32 {
	return in.readWord(intSyscallPC)
}

// Issue issues an interrupt. If the interrupt is already issued,
// this has no effect.
func (in *Interrupt) Issue(i byte) {
	off := uint32(i/8) + intPending
	b := in.readByte(off)
	b |= 0x1 << (i % 8)
	in.writeByte(off, b)
}

// Clear clears an interrupt. If the interrupt is already cleared,
// this has no effect.
func (in *Interrupt) Clear(i byte) {
	off := uint32(i/8) + intPending
	b := in.readByte(off)
	b &= ^(0x1 << (i % 8))
	in.writeByte(off, b)
}

// Enable sets the interrupt enable bit in the flags.
func (in *Interrupt) Enable() {
	b := in.readByte(intFlags)
	b |= 0x1
	in.writeByte(intFlags, b)
}

// Enabled tests if interrupt is enabled
func (in *Interrupt) Enabled() bool {
	b := in.readByte(intFlags)
	return (b & 0x1) != 0
}

// Disable clears the interrupt enable bit in the flags.
func (in *Interrupt) Disable() {
	b := in.readByte(intFlags)
	b &= ^byte(0x1)
	in.writeByte(intFlags, b)
}

// EnableInt enables a particular interrupt by clearing the mask.
func (in *Interrupt) EnableInt(i byte) {
	off := uint32(i/8) + intMask
	b := in.readByte(off)
	b |= 0x1 << (i % 8)
	in.writeByte(off, b)
}

// DisableInt disables a particular interrupt by setting the mask.
func (in *Interrupt) DisableInt(i byte) {
	off := uint32(i/8) + intMask
	b := in.readByte(off)
	b &= ^(0x1 << (i % 8))
	in.writeByte(off, b)
}

// Flags returns the flags byte.
func (in *Interrupt) Flags() byte {
	return in.readByte(intFlags)
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
		pending := in.readByte(intPending + i)
		mask := in.readByte(intMask + i)
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
