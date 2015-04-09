package g8

import (
	"fmt"

	"lonnie.io/e8vm/g8/ast"
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
	*ref
}

type objFunc struct {
	name string
	*ref
	f *ast.Func
}

type objConst struct {
	name string
	*ref
}

type objType struct {
	name string
	*ref
}
