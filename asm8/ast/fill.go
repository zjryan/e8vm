package ast

// Filling methods for instructions
const (
	FillNone = iota
	FillLink // for jumps
	FillLow  // for immediate instructions
	FillHigh // for lui
	FillLabel
)
