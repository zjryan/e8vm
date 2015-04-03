package build8

import (
	"io"
	"path/filepath"
)

// SingleFile creates a file map that can be used for single file compilation
func SingleFile(path string, rc io.ReadCloser) map[string]*File {
	name := filepath.Base(path)
	file := &File{
		Name:       name,
		Path:       path,
		ReadCloser: rc,
	}
	return map[string]*File{name: file}
}

// SimplePkg creates a package that contains only one file
// and has no imports
func SimplePkg(p string, f string, rc io.ReadCloser) *PkgInfo {
	single := SingleFile(f, rc)
	return &PkgInfo{
		Path: p,
		Src:  single,
	}
}
