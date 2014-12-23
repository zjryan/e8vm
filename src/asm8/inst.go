package asm8

type inst struct {
	inst   uint32
	pack   string
	symbol string
	fill   int
}

const (
	fillNone = iota
	fillLabel
	fillLink
	fillLow
	fillHigh
)
