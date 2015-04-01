package g8

import (
	"fmt"

	"lonnie.io/e8vm/g8/ir"
)

func addressable(r ir.Ref) bool {
	// TODO:
	return true
}

func canAssign(left, right typ) bool {
	return sameType(left, right)
}

func sameType(t1, t2 typ) bool {
	if t1 == t2 {
		return true
	}

	switch t1 := t1.(type) {
	case typBasic:
		if t2, ok := t2.(typBasic); ok {
			return t2 == t1
		}
		return false
	case *typPtr:
		if t2, ok := t2.(*typPtr); ok {
			return sameType(t1.t, t2.t)
		}
		return false
	case *typSlice:
		if t2, ok := t2.(*typSlice); ok {
			return sameType(t1.t, t2.t)
		}
		return false
	case *typArray:
		if t2, ok := t2.(*typArray); ok {
			return t1.n == t2.n && sameType(t1.t, t2.t)
		}
		return false
	case *typFunc:
		if t2, ok := t2.(*typFunc); ok {
			if len(t1.argTypes) != len(t2.argTypes) {
				return false
			}
			if len(t1.retTypes) != len(t2.argTypes) {
				return false
			}

			for i, t := range t1.argTypes {
				if !sameType(t, t2.argTypes[i]) {
					return false
				}
			}

			for i, t := range t2.retTypes {
				if !sameType(t, t2.retTypes[i]) {
					return false
				}
			}

			return true
		}
		return false
	case *typStruct:
		panic("todo") // not clear what to do here yet...
	default:
		panic(fmt.Errorf("invalid type: %t", t1))
	}
}
