package asm8

// asm8 symbol types
const (
	SymImport = iota
	SymFunc
	SymConst
	SymVar
	SymLabel
)

func symStr(s int) string {
	switch s {
	case SymImport:
		return "import"
	case SymFunc:
		return "function"
	case SymConst:
		return "constant"
	case SymVar:
		return "variable"
	case SymLabel:
		return "label"
	}
	return "unknown"
}
