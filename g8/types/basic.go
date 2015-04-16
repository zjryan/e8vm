package types

import (
	"fmt"
)

// Basic types are the built-in fundamentals of the language
type Basic int

// Basic type codes
const (
	Int Basic = iota
	Uint
	Int8
	Uint8
	Float32
	Bool
	String
)

// IsBasic checks if a type is a particular basic type
func IsBasic(t T, b Basic) bool {
	code, ok := t.(Basic)
	if !ok {
		return false
	}
	return code == b
}

// IsByte checks if a type is Uint8
func IsByte(t T) bool {
	code, ok := t.(Basic)
	if !ok {
		return false
	}
	return code == Uint8
}

// BothBasic checks if two types are both a particular basic type
func BothBasic(a, b T, t Basic) bool {
	return IsBasic(a, t) && IsBasic(b, t)
}

// Size returns the size in memory of a basic type
func (t Basic) Size() int32 {
	switch t {
	case Int, Uint:
		return 4
	case Int8, Uint8:
		return 1
	case Float32:
		return 4
	case String:
		panic("todo")
	case Bool:
		return 1
	default:
		panic("unknown basic type")
	}
}

// String returns the name of the basic type.
func (t Basic) String() string {
	switch t {
	case Int:
		return "int"
	case Uint:
		return "uint"
	case Int8:
		return "int8"
	case Uint8:
		return "uint8"
	case Bool:
		return "bool"
	case String:
		return "string"
	default:
		panic(fmt.Errorf("invalid basic type %d", t))
	}
}
