package g8

import (
	"fmt"
)

type typBasic int

const (
	typVoid typBasic = iota
	typInt
	typUint
	typInt8
	typUint8
	typFloat32
	typBool
	typString
)

func isVoid(a typ) bool { return isBasic(a, typVoid) }

func isBasic(a typ, t typBasic) bool {
	code, ok := a.(typBasic)
	if !ok {
		return false
	}
	return code == t
}

func bothBasic(a, b typ, t typBasic) bool {
	return isBasic(a, t) && isBasic(b, t)
}

func (t typBasic) Size() int32 {
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
}

func (t typBasic) String() string {
	switch t {
	case typVoid:
		return "void"
	case typInt:
		return "int"
	case typUint:
		return "uint"
	case typInt8:
		return "int8"
	case typUint8:
		return "uint8"
	case typBool:
		return "bool"
	case typString:
		return "string"
	default:
		panic(fmt.Errorf("invalid basic type %d", t))
	}
}
