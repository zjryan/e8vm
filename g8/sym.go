package g8

import (
	"fmt"
)

const (
	symNone = iota
	symFunc
	symVar
	symType
	symConst
	symImport
)

func symStr(s int) string {
	switch s {
	case symVar:
		return "variable"
	case symType:
		return "type"
	case symFunc:
		return "function"
	case symConst:
		return "constant"
	case symImport:
		return "imported package"
	default:
		panic(fmt.Errorf("unknown symbol: %d", s))
	}
}

type objVar struct {
	name string
	*ref // the reference of this variable
}

type objFunc struct {
	name string
	*ref
}

type objConst struct {
	name string
	*ref
}

type objType struct {
	name string
	*ref
}
