package g8

import (
	"lonnie.io/e8vm/lex8"
)

// CompileSingleFile compiles a file into a bare-metal E8 image
func CompileSingleFile(fname, s string) ([]byte, []*lex8.Error) {
	return buildSingle(fname, s, Lang())
}
