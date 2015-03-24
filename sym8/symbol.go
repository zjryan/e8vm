package sym8

import (
	"lonnie.io/e8vm/lex8"
)

// Symbol is a data structure for saving a symbol.
type Symbol struct {
	Name string
	Type int
	Item interface{}
	Pos  *lex8.Pos
}

// Clone deep copies the symbol
func (s *Symbol) Clone() *Symbol {
	return &Symbol{s.Name, s.Type, s.Item, s.Pos}
}

// Make creates a new symbol
func Make(name string, t int, item interface{}, pos *lex8.Pos) *Symbol {
	return &Symbol{name, t, item, pos}
}
