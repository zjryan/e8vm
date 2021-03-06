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
	Path     string
	Pos      *lex8.Pos
	Compiled Linkable
}

// Linkable is an interface for a linkable package
type Linkable interface {
	Main() string
	Lib() *link8.Pkg
}

// Importer is an interface for importing required packages for compiling
type Importer interface {
	Import(name, path string, pos *lex8.Pos) // imports a package
}

// PkgInfo contains the information for compiling a package
type PkgInfo struct {
	Path   string
	Src    map[string]*File
	Import map[string]*Import

	CreateLog func(name string) io.WriteCloser
}

// Lang is a language compiler interface
type Lang interface {
	// IsSrc filters source file filenames
	IsSrc(filename string) bool

	// Prepare issues import requests
	Prepare(src map[string]*File, importer Importer) []*lex8.Error

	// Compile compiles a list of source files into a compiled linkable
	Compile(pinfo *PkgInfo) (Linkable, []*lex8.Error)
}
