package lex8

// Token defines a token structure.
type Token struct {
	Type int
	Lit  string
	Pos  *Pos
}

const (
	EOF     = 0
	Illegal = -1
)
