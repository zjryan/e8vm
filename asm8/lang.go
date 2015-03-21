package asm8

import (
	"io"
	"path"
	"strings"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/pkg8"
)

type lang struct{}

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".s")
}

func (lang) ListImport(src pkg8.Files) ([]string, []*lex8.Error) {
	for f, rc := range src {
		if path.Base(f) == "import.s" || len(src) == 1 {
			return listImport(f, rc)
		}
	}
	return nil, nil
}

func (lang) Compile(
	path string,
	src pkg8.Files,
	importer pkg8.Importer,
) (
	lib pkg8.Linkable,
	errs []*lex8.Error,
) {
	panic("todo")
}

func (lang) Load(r io.Reader) error {
	panic("todo")
}

var Lang pkg8.Lang = lang{}
