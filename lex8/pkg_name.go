package lex8

// IsPkgName checks if a literal is a valid package name.
func IsPkgName(s string) bool {
	// must not an empty string
	if s == "" {
		return false
	}

	// must not a keyword
	if s == "func" || s == "var" || s == "const" {
		return false
	}

	for i, r := range s {
		if i > 0 && r >= '0' && r <= '9' {
			if i > 0 {
				continue
			}
			return false
		}

		if r >= 'a' && r <= 'z' {
			continue
		}
		return false
	}
	return true
}
