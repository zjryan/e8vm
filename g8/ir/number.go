package ir

import (
	"fmt"
)

type number struct{ v uint32 } // a constant number

func (n *number) String() string { return fmt.Sprintf("%d", n.v) }

// Num creates a constant reference to a int32 number
func Num(v uint32) Ref { return &number{v} }

// Snum creates a constant reference to a uint32 number
func Snum(v int32) Ref { return &number{uint32(v)} }
