package asm8

// File represents a file.
type file struct {
	Funcs []*funcDecl
	Vars  []*varDecl
}
