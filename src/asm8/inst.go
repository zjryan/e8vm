package asm8

import (
	"lex8"
)

type inst struct {
	inst   uint32
	pack   string
	symbol string
	fill   int

	extras []uint32 // for pseudo asms

	symTok *lex8.Token
}

const (
	fillNone = iota
	fillLabel
	fillLink
	fillLow
	fillHigh
)

func isJump(inst uint32) bool {
	return (inst >> 31) > 0
}

func inBrRange(delta uint32) bool {
	d := int32(delta)
	return d >= -0x20000 && d <= 0x1ffff
}
