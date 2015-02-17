package build8

import (
	"strings"
)

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

func isPkgPath(p string) bool {
	if p == "" {
		return false
	}
	subs := strings.Split(p, "/")
	for _, sub := range subs {
		if !isPkgName(sub) {
			return false
		}
	}
	return true
}
