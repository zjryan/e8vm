package ast

// Package represents a package node.
type Pkg struct {
	Path    string // package import path
	Files   []*File
	Imports map[string]*PkgImport
}

// NewPackage creates an empty package node.
func NewPkg(path string) *Pkg {
	ret := new(Pkg)
	ret.Path = path
	return ret
}

// AddFile adds a file into the package.
func (p *Pkg) AddFile(f *File) {
	p.Files = append(p.Files, f)
}
