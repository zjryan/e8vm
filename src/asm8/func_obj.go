package asm8

// FuncObj is a function object that is ready for linking.
type FuncObj struct {
}

// InstLine what we map from an instruction to in an object file
type instLine struct {
	funcOffset uint32 // offset in function
	inst       uint32 // instruction prototype

	fill     func(line *instLine) // filling methods
	symbol   string               // the symbol
	symValue uint32               // symbol value
}
