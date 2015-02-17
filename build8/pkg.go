package build8

func isPkgName(s string) bool {
	if s == "" {
		return false
	}

	for i, r := range s {
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		return false
	}

	return true
}
