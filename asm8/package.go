package asm8

// Package represents a package node.
type Package struct {
	Path  string
	Files []*File
}

// NewPackage creates an empty package node.
func NewPackage(path string) *Package {
	ret := new(Package)
	ret.Path = path
	return ret
}

// AddFile adds a file into the package.
func (p *Package) AddFile(f *File) {
	p.Files = append(p.Files, f)
}
