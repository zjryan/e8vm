package lex8

// Token defines a token structure.
type Token struct {
	Type int
	Lit  string
	Pos  *Pos
}

// Standard token types
const (
	EOF     = -1
	Illegal = -2
)
