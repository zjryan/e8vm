package asm8

type asmLine struct {
	inst  uint32
	label string
	fill  int // the filing method
}
