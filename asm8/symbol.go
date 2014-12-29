package asm8

import (
	"lonnie.io/e8vm/lex8"
)

// Symbol is a data structure for saving a symbol.
type Symbol struct {
	Name    string
	Type    int
	Item    interface{}
	Pos     *lex8.Pos
	Package string // Package path
}

func (s *Symbol) clone() *Symbol {
	return &Symbol{s.Name, s.Type, s.Item, s.Pos, s.Package}
}

// asm8 symbol types
const (
	SymImport = iota // Item.type == *PkgObj
	SymFunc			 // Item.type == *Func
	SymConst		 // Item.type == *Const	
	SymVar			 // Item.type == *Var
	SymLabel		 // Item.type == *stmt
)

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

// IsPublic checks if a symbol name is public.
func IsPublic(name string) bool {
	if name == "" {
		return false
	}
	r := name[0]
	return r >= 'A' && r <= 'Z'
}
