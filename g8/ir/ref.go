package ir

import (
	"fmt"
)

// Ref is a reference of an object
type Ref interface{}

// stackVar is a variable on stack
type stackVar struct {
	name   string // not unique, just for debugging
	id     int
	offset int32
	size   int32

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

func (v *stackVar) String() string {
	if v.name != "" {
		return v.name
	}
	return fmt.Sprintf("<%d>", v.id)
}

type heapSym struct{ pkg, sym int } // a variable on heap
type funcSym struct{ pkg, sym int } // a function symbol
type number struct{ v uint32 }      // a constant number

func (n *number) String() string { return fmt.Sprintf("%d", n.v) }

// Num creates a constant reference to a int32 number
func Num(v uint32) Ref { return &number{v} }

// Snum creates a constant reference to a uint32 number
func Snum(v int32) Ref { return &number{uint32(v)} }
