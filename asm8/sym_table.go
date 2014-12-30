package asm8

// SymTable save the symbol
type symTabel struct {
	m map[string]*symbol
}

// NewSymTable creates an empty symbol table
func newSymTable() *symTabel {
	ret := new(symTabel)
	ret.m = make(map[string]*symbol)

	return ret
}

// Query searches for a symbol with a particular name.
func (tab *symTabel) Query(n string) *symbol {
	s := tab.m[n]
	if s == nil {
		return nil
	}

	return s.clone()
}

// Declare adds a symbol into the table.
// It returns nil on successful, and returns the conflict symbol
// when it is already declared.
func (tab *symTabel) Declare(s *symbol) *symbol {
	n := s.Name
	p := tab.m[n]
	if p != nil {
		return p.clone()
	}

	tab.m[n] = s.clone()
	return nil
}
