package asm8

import (
	"io"
	"strings"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/pkg8"
)

type lang struct{}

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".s")
}

func (lang) ListImport(src pkg8.Files) ([]string, []*lex8.Error) {
	if len(src) == 1 {
		panic("todo")
	}

	imp := src["import.s"]
	if imp == nil {
		return nil, nil
	}

	panic("todo")
}

func (lang) Compile(
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
