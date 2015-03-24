package sym8

// Scope is a stack of symbol tables.
type Scope struct {
	stack []*Table
	top   *Table // the top
}

// NewScope creates a new symbole scope with one symbol table at the
// bottom.
func NewScope() *Scope {
	ret := new(Scope)
	ret.Push()
	return ret
}

// Push adds a new symbol table on the top of the stack.
func (s *Scope) Push() {
	tab := NewTable()
	s.top = tab
	s.stack = append(s.stack, tab)
}

// Pop removes a symbol table from the top of the stack.
// It panics when the stack is empty after the pop.
func (s *Scope) Pop() {
	n := len(s.stack)
	if n < 2 {
		panic("stack empty after pop")
	}

	s.stack = s.stack[:n-1]
	s.top = s.stack[n-2]
}

// Depth returns the number of scopes on the stack.
func (s *Scope) Depth() int {
	return len(s.stack)
}

// Query searches for the top visible symbol with a particular name.
func (s *Scope) Query(n string) *Symbol {
	d := len(s.stack)

	for i := d - 1; i >= 0; i-- {
		tab := s.stack[i]
		ret := tab.Query(n)
		if ret != nil {
			return ret
		}
	}

	return nil
}

// Declare declares a symbole on the top of the symbol table stack.
// It returns nil on successful, and returns the conflict symbol when
// it is already declared.
func (s *Scope) Declare(sym *Symbol) *Symbol {
	return s.top.Declare(sym)
}
