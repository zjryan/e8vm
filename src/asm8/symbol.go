package asm8

import (
	"lex8"
)

// Symbol is a data structure for saving a symbol.
type Symbol struct {
	Name string
	Type int
	Item interface{}
	Pos  *lex8.Pos
}

func (s *Symbol) clone() *Symbol {
	return &Symbol{s.Name, s.Type, s.Item, s.Pos}
}

// asm8 symbol types
const (
	SymImport = iota
	SymFunc
	SymConst
	SymVar
	SymLabel
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
