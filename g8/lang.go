package g8

import (
	"strings"

	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/lex8"
)

type lang struct{}

// Lang returns the G language builder for the building system
func Lang() build8.Lang { return lang{} }

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".g")
}

func (lang) Import(p build8.Pkg) []*lex8.Error {
	p.AddImport("$", "asm/builtin", nil)
	// TODO: parse import statements
	return nil
}

func (lang) Compile(p build8.Pkg) []*lex8.Error {
	return nil
}
