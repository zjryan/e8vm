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

func (lang) Prepare(
	src map[string]*build8.File, importer build8.Importer,
) []*lex8.Error {
	importer.Import("$", "asm/builtin", nil)
	// TODO: parse import statements
	return nil
}

func (lang) Compile(
	path string, src map[string]*build8.File, imp map[string]*build8.Import,
) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	panic("todo")
}
