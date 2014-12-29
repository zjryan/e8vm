package asm8

import (
	"lonnie.io/e8vm/lex8"
)

type inst struct {
	inst   uint32
	pack   string
	symbol string
	fill   int

	symTok *lex8.Token
}

const (
	fillNone = iota
	fillLow  // for immediate instructions
	fillHigh // for lui
	fillLink // for jumps
	fillLabel
)

func isJump(inst uint32) bool {
	return (inst >> 31) > 0
}

func inBrRange(delta uint32) bool {
	d := int32(delta)
	return d >= -0x20000 && d <= 0x1ffff
}
