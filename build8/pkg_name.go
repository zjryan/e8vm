package build8

import (
	"strings"

	"lonnie.io/e8vm/lex8"
)

func isPkgPath(p string) bool {
	if p == "" {
		return false
	}
	subs := strings.Split(p, "/")
	for _, sub := range subs {
		if !lex8.IsPkgName(sub) {
			return false
		}
	}
	return true
}
