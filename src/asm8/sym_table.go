package asm8

import (
	"errors"
)

type symbol struct {
	Name string
	Type int
	Item interface{}
}

// SymTable save the symbol
type SymTable struct {
	m    map[string]*symbol
	lsts map[int][]*symbol
}

// NewSymTable creates an empty symbol table
func NewSymTable() *SymTable {
	ret := new(SymTable)
	ret.m = make(map[string]*symbol)
	ret.lsts = make(map[int][]*symbol)

	return ret
}

var errSymExists = errors.New("symbol exists")

// Query look for a symbol with a particular name.
func (tab *SymTable) Query(n string) (i interface{}, t int) {
	s := tab.m[n]
	if s == nil {
		return nil, 0
	}

	return s.Item, s.Type
}

// Declare adds a symbol into the table.
func (tab *SymTable) Declare(n string, t int, i interface{}) error {
	if tab.m[n] != nil {
		return errSymExists
	}

	s := &symbol{n, t, i}
	tab.m[n] = s
	tab.lsts[t] = append(tab.lsts[t], s)
	return nil
}

// List returns the list of symbols of a particular type.
func (tab *SymTable) List(t int) []interface{} {
	var ret []interface{}

	lst := tab.lsts[t]
	for _, s := range lst {
		ret = append(ret, s.Item)
	}

	return ret
}
