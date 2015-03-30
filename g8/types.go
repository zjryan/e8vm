package g8

import (
	"fmt"
)

type typ interface{}

type typBasic int

const (
	typErr typBasic = iota
	typInt
	typUint
	typInt8
	typUint8
	typFloat32
	typBool
	typString
)

type typPtr struct{ t typ } // a pointer type
type typSlice struct{ t typ }
type typArray struct {
	t typ
	n uint32
}

type varSig struct {
	name string
	t    typ
}

type funcSig struct {
	name string
	t    *typFunc
}

type typStruct struct {
	fields  []*varSig
	methods []*funcSig
}

type typFunc struct {
	argTypes []typ
	regTypes []typ

	// optional names
	argNames []string
	retNames []string
}

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
