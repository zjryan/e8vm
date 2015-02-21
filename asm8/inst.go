package asm8

const (
	fillNone = iota
	fillLink // for jumps
	fillLow  // for immediate instructions
	fillHigh // for lui
	fillLabel
)

func isJump(inst uint32) bool {
	return (inst >> 31) > 0
}

func inBrRange(delta uint32) bool {
	d := int32(delta)
	return d >= -0x20000 && d <= 0x1ffff
}
