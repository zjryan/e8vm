package ir

import (
	"fmt"
)

// stackVar is a variable on stack
type stackVar struct {
	name string // not unique, just for debugging
	id   int

	// the offset relative to SP
	// before SP shift, the variable is saved at [SP-offset]
	// after SP shift, the variable is saved at [SP+framesize-offset]
	offset int32
	size   int32
	u8     bool

	// reg is the register allocated
	// valid values are in range [1, 4] for normal values
	// and also ret register is 6
	viaReg uint32

	// regOnly stack vars does not take frame space on the stack
	regOnly bool
}

func newVar(n int32, name string, isByte bool) *stackVar {
	ret := new(stackVar)
	ret.name = name
	ret.size = n
	ret.u8 = isByte

	return ret
}

func (v *stackVar) String() string {
	if v.name != "" {
		return v.name
	}
	return fmt.Sprintf("<%d>", v.id)
}

func (v *stackVar) Addressable() bool { return true }
