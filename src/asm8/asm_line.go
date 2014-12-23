package asm8

type asmLine struct {
	ops []*Token

	inst  uint32
	label string
	fill  int // the filing method
}
