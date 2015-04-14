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

func newVar(n int32, name string) *stackVar {
	ret := new(stackVar)
	ret.name = name
	ret.size = n

	return ret
}

func newByte(name string) *stackVar {
	ret := newVar(1, name)
	ret.u8 = true
	return ret
}

func (v *stackVar) String() string {
	if v.name != "" {
		return v.name
	}
	return fmt.Sprintf("<%d>", v.id)
}
