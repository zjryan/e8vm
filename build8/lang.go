package build8

import (
	"io"

	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// File is a file in a package.
type File struct {
	Name string
	Path string
	io.ReadCloser
}

// Import is an import identity
type Import struct {
	Path string
	Pos  *lex8.Pos

	Pkg Pkg // filled by the build system
}

// Linkable is an interface for a linkable package
type Linkable interface {
	Lib() *link8.Pkg
}

// Pkg is a package interface for interracting with the language
type Pkg interface {
	Path() string
	Src() map[string]*File

	AddImport(name, path string, pos *lex8.Pos)
	Imports() map[string]*Import

	SetCompiled(lib Linkable)
	Compiled() Linkable
}

// Lang is a language compiler interface
type Lang interface {
	IsSrc(filename string) bool
	Import(p Pkg) []*lex8.Error
	Compile(p Pkg) []*lex8.Error
}
