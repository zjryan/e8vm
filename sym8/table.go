package sym8

// Table save the symbol
type Table struct {
	m map[string]*Symbol
}

// NewTable creates an empty symbol table
func NewTable() *Table {
	ret := new(Table)
	ret.m = make(map[string]*Symbol)
	return ret
}

// Query searches for a symbol with a particular name.
func (tab *Table) Query(n string) *Symbol {
	s := tab.m[n]
	if s == nil {
		return nil
	}

	return s.Clone()
}

// Declare adds a symbol into the table.
// It returns nil on successful, and returns the conflict symbol
// when it is already declared.
func (tab *Table) Declare(s *Symbol) *Symbol {
	n := s.Name
	p := tab.m[n]
	if p != nil {
		return p.Clone()
	}

	tab.m[n] = s.Clone()
	return nil
}
