package g8

import (
	"fmt"
)

func typSize(t typ) int32 {
	switch t := t.(type) {
	case typBasic:
		switch t {
		case typVoid:
			return 0
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
	case *typPtr:
		return 4
	case *typSlice:
		panic("todo")
	case *typFunc:
		return 4
	case *typStruct:
		panic("todo")
	default:
		panic(fmt.Errorf("invalid type: %T", t))
	}
}
