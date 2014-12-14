package asm8

// Line is an assembly line.
type Line struct {
	Ops []*Token
}

// Func is an assembly function.
type Func struct {
	Lines []*Line

	kw     *Token
	name   *Token
	lbrace *Token
	rbrace *Token
}
