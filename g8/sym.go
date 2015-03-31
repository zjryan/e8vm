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
