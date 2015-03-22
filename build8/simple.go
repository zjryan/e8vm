package build8

import (
	"io"
	"path/filepath"

	"lonnie.io/e8vm/lex8"
)

type simple struct {
	path string
	io.ReadCloser
	lib Linkable
}

// NewSimplePkg creates a single-file stand-alone package.
func NewSimplePkg(path string, rc io.ReadCloser) Pkg {
	ret := new(simple)
	ret.path = path
	ret.ReadCloser = rc

	return ret
}

func (s *simple) Src() map[string]*File {
	ret := make(map[string]*File)
	name := filepath.Base(s.path)
	ret[name] = &File{
		Name:       name,
		Path:       s.path,
		ReadCloser: s.ReadCloser,
	}
	return ret
}

func (s *simple) AddImport(name, path string, pos *lex8.Pos) {
	panic("not impl.")
}

func (s *simple) Path() string                { return "_" }
func (s *simple) Imports() map[string]*Import { return nil }
func (s *simple) SetCompiled(lib Linkable)    { s.lib = lib }
func (s *simple) Compiled() Linkable          { return s.lib }
