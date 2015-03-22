package asm8

import (
	"io"
	"path"
	"strings"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/pkg8"
)

type lang struct{}

func (lang) IsSrc(filename string) bool {
	return strings.HasSuffix(filename, ".s")
}

func (lang) ListImport(src pkg8.Files) ([]string, []*lex8.Error) {
	for f, rc := range src {
		if len(src) == 1 || path.Base(f) == "import.s" {
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
	pkg8.Linkable,
	[]*lex8.Error,
) {
	pkg := new(pkg)
	pkg.path = path

	errs := lex8.NewErrorList()

	for f, rc := range src {
		astFile, es := parse.File(f, rc)
		if es != nil {
			return nil, es
		}

		file := resolveFile(errs, astFile)
		if len(errs.Errs) != 0 {
			return nil, errs.Errs
		}

		pkg.files = append(pkg.files, file)
	}

	b := newBuilder()
	lib := buildLib(b, pkg)
	es := b.Errs()
	if len(es) != 0 {
		return nil, es
	}

	return lib, nil
}

func (lang) Load(r io.Reader) error {
	panic("todo")
}

var Lang pkg8.Lang = lang{}
