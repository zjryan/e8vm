package asm8

type SymScope struct {
	stack []*SymTable
	top   *SymTable // the top
}

func NewSymScope() *SymScope {
	ret := new(SymScope)
	ret.Push()

	return ret
}

// Push adds a new symbol table on the top of the stack.
func (s *SymScope) Push() {
	tab := NewSymTable()
	s.top = tab
	s.stack = append(s.stack, tab)
}

// Pop removes a symbol table from the top of the stack.
// It panics when the stack is empty after the pop.
func (s *SymScope) Pop() {
	n := len(s.stack)
	if n < 2 {
		panic("stack empty after pop")
	}

	s.stack = s.stack[:n-1]
	s.top = s.stack[n-2]
}

// Depth returns the number of scopes on the stack.
func (s *SymScope) Depth() int {
	return len(s.stack)
}

// Query searches for the top visible symbol with a particular name.
func (s *SymScope) Query(n string) (item interface{}, t int) {
	d := len(s.stack)

	for i := d - 1; i >= 0; i-- {
		tab := s.stack[i]
		item, t = tab.Query(n)
		if item != nil {
			return item, t
		}
	}

	return nil, 0
}

// Declare declares a symbole on the top of the symbol table stack.
func (s *SymScope) Declare(n string, t int, i interface{}) error {
	return s.top.Declare(n, t, i)
}
