package asm8

// asm8 token types.
const (
	EOF = iota
	Comment
	Keyword
	Operand
	String
	Lbrace
	Rbrace
	Endl
	Illegal
)

// Token defines a token structure.
type Token struct {
	Type int
	Lit  string
	Pos  *Pos
}
