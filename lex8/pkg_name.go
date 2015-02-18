package lex8

// IsPkgName checks if a literal is a valid package name.
func IsPkgName(s string) bool {
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
