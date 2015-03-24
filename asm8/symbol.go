package asm8

import (
	"lonnie.io/e8vm/link8"
)

// Symbol types
const (
	SymNone   = iota
	SymFunc   // Item.type == *Func
	SymVar    // Item.type == *Var
	SymConst  // Item.type == *Const // TODO
	SymImport // Item.type == *PkgImport
	SymLabel  // Item.type == *stmt
)

func init() {
	as := func(b bool) {
		if !b {
			panic("bug")
		}
	}
	as(SymNone == link8.SymNone)
	as(SymFunc == link8.SymFunc)
	as(SymVar == link8.SymVar)
}

func symStr(s int) string {
	switch s {
	case SymImport:
		return "import"
	case SymFunc:
		return "function"
	case SymConst:
		return "constant"
	case SymVar:
		return "variable"
	case SymLabel:
		return "label"
	}
	return "unknown"
}
