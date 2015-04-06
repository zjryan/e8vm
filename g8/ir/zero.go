package ir

import (
	"fmt"
)

func zeroRef(b *Block, r Ref) {
	switch r := r.(type) {
	case *stackVar:
		saveVar(b, 0, r)
	case *number:
		panic("number are read only")
	default:
		panic(fmt.Errorf("not implemented: %T", r))
	}
}
