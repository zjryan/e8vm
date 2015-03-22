package pkg8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Linkable is an interface for a linkable package
type Linkable interface {
	Lib() *link8.Pkg
	Save(w io.Writer) error
}

// Importer imports a package
type Importer interface {
	Import(path string) Linkable
}

// Files is just a map from file names to file read/closers.
type Files map[string]io.ReadCloser

// Lang is a language compiler
type Lang interface {
	IsSrc(filename string) bool
	ListImport(src Files) ([]string, []*lex8.Error)
	Compile(p string, src Files, imp Importer) (Linkable, []*lex8.Error)
	Load(r io.Reader) error
}