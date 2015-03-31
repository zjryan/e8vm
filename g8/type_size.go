package g8

import (
	"fmt"
)

func typeSize(t typ) int32 {
	switch t := t.(type) {
	case typBasic:
		switch t {
		case typErr:
			return 4
		case typInt, typUint:
			return 4
		case typInt8, typUint8:
			return 1
		case typFloat32:
			return 4
		case typString:
			panic("todo")
		default:
			panic("unknown basic type")
		}
	case typPtr:
		return 4
	case typSlice:
		panic("todo")
	case typFunc:
		return 4
	default:
		panic(fmt.Errorf("unknown type: %T", t))
	}
}