package build8

import (
	"io"
)

// Home is a storage place for storing building files
type Home interface {
	// Pkgs lists all the packages
	Pkgs(prefix string) []string

	// Src lists the source files in a package
	Src(path string) map[string]*File

	// CreateLib creates the writer for writing the linkable package library
	CreateLib(path string) io.WriteCloser

	// CreateLog creates a logger, usually for debugging
	CreateLog(path, name string) io.WriteCloser

	// Bin creates the writer for generate the E8 binary
	CreateBin(path string) io.WriteCloser

	// Lang returns the language of a path
	Lang(path string) Lang
}
