package link8

// Symbol is a linking symbol in an object file.
type Symbol struct {
	Name string
	Type int
}

// Linking symbol types
const (
	SymFunc = iota
	SymVar
)
