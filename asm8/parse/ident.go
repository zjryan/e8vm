package parse

// IsIdent checks if a string is a valid identifier
func IsIdent(id string) bool {
	if id == "" {
		return false
	}
	for i, r := range id {
		if r >= '0' && r <= '9' {
			if i > 0 {
				continue
			}
			return false
		}

		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r == '_' || r == ':' {
			continue
		}
		return false
	}
	return true
}
