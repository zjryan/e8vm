package build8

import (
	"strings"
)

const importFile = "imports"

func isImports(base string) bool {
	return base == importFile
}

func isAsm(base string) bool {
	return strings.HasSuffix(base, ".s")
}

func isSrc(base string) bool {
	return isImports(base) || isAsm(base)
}
