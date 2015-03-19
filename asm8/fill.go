package asm8

// Filling methods for instructions
const (
	fillNone = iota
	fillLink // for jumps
	fillLow  // for immediate instructions
	fillHigh // for lui
	fillLabel
)
