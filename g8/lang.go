package g8

import (
	"fmt"
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

func initBuilder(b *builder, imp map[string]*build8.Import) {
	builtin, ok := imp["$"]
	if !ok {
		b.Errorf(nil, "builtin import missing for %q", b.path)
		return
	}

	declareBuiltin(b, builtin.Compiled.Lib())
}

func (lang) Compile(pinfo *build8.PkgInfo) (
	compiled build8.Linkable, es []*lex8.Error,
) {
	// need to load these two builtin functions here
	b := newBuilder(pinfo.Path)
	initBuilder(b, pinfo.Import)

	if es = b.Errs(); es != nil {
		return nil, es
	}

	return nil, lex8.SingleErr(fmt.Errorf("g8 compiler not complete"))
}
