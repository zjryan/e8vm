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
	m map[string]*symbol
}

// NewSymTable creates an empty symbol table
func NewSymTable() *SymTable {
	ret := new(SymTable)
	ret.m = make(map[string]*symbol)

	return ret
}

var errSymExists = errors.New("symbol exists")

// Query searches for a symbol with a particular name.
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
	return nil
}
