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
