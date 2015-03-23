package parse

func isKeyword(lit string) bool {
	switch lit {
	case "func", "var", "const", "import", "type",
		"if", "else", "for",
		"break", "continue", "return",
		"switch", "case", "default", "fallthrough",
		"range", "struct", "interface", "goto":
		return true
	}
	return false
}
