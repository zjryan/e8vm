package lex8

import (
	"io"
)

// Tokener is token emitting interface.
type Tokener interface {
	// Token() returns the next token
	Token() *Token

	// Errs returns the error list on tokening
	Errs() []*Error
}

// NewTokener creates a new tokener from LexFunc x and WhiteFunc w.
func NewTokener(f string, r io.Reader, x LexFunc, w WhiteFunc) Tokener {
	ret := NewLexer(f, r)
	ret.LexFunc = x
	ret.IsWhite = w

	return ret
}
