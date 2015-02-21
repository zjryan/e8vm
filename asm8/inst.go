package asm8

func isJump(inst uint32) bool {
	return (inst >> 31) > 0
}

func inBrRange(delta uint32) bool {
	d := int32(delta)
	return d >= -0x20000 && d <= 0x1ffff
}
