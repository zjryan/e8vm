package types

import (
	"fmt"
)

// CanAssign checks if right can be assigned to right
func CanAssign(left, right Type) bool {
	return SameType(left, right)
}

// SameType checks if two types are of the same type
func SameType(t1, t2 Type) bool {
	if t1 == t2 {
		return true
	}

	switch t1 := t1.(type) {
	case Basic:
		if t2, ok := t2.(Basic); ok {
			return t2 == t1
		}
		return false
	case *Pointer:
		if t2, ok := t2.(*Pointer); ok {
			return SameType(t1.T, t2.T)
		}
		return false
	case *Slice:
		if t2, ok := t2.(*Slice); ok {
			return SameType(t1.T, t2.T)
		}
		return false
	case *Array:
		if t2, ok := t2.(*Array); ok {
			return t1.N == t2.N && SameType(t1.T, t2.T)
		}
		return false
	case *Func:
		if t2, ok := t2.(*Func); ok {
			if len(t1.Args) != len(t2.Args) {
				return false
			}
			if len(t1.Rets) != len(t2.Rets) {
				return false
			}

			for i, t := range t1.Args {
				if !SameType(t.Type, t2.Args[i].Type) {
					return false
				}
			}

			for i, t := range t2.Rets {
				if !SameType(t.Type, t2.Rets[i].Type) {
					return false
				}
			}

			return true
		}
		return false
	case *Struct:
		panic("todo") // not clear what to do here yet...
	default:
		panic(fmt.Errorf("invalid type: %t", t1))
	}
}
