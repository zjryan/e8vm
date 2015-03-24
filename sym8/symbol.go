package sym8

import (
	"lonnie.io/e8vm/lex8"
)

// Symbol is a data structure for saving a symbol.
type Symbol struct {
	name string

	Type int
	Item interface{}
	Pos  *lex8.Pos
}

// Name returns the symbol name.
// This name is immutable for its used for indexing in the tables.
func (s *Symbol) Name() string { return s.name }

// Make creates a new symbol
func Make(name string, t int, item interface{}, pos *lex8.Pos) *Symbol {
	return &Symbol{name, t, item, pos}
}
