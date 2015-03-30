package link8

// Symbol is a linking symbol in an object file
type Symbol struct {
	Name string
	Type int
}

// Linking symbol types
const (
	SymNone = iota // for default return value
	SymFunc
	SymVar
)

func symStr(t int) string {
	switch t {
	case SymNone:
		return "none"
	case SymFunc:
		return "func"
	case SymVar:
		return "var"
	default:
		return "unknown"
	}
}
