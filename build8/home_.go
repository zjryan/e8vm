package build8

import (
	"io"
)

// Home is a storage place for storing building files
type Home interface {
	// Pkgs lists all the packages
	Pkgs() []string

	// Src lists the source files in a package
	Src(path string) (map[string]*File, Lang)

	// Lib creates the writer for writing the linkable package library
	Lib(path string) io.WriteCloser

	// Log creates a logger, usually for debugging
	Log(path, name string) io.WriteCloser

	// Bin creates the writer for generate the E8 binary
	Bin(path string) io.WriteCloser
}
