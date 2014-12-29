package asm8

// File represents a file.
type File struct {
	Funcs []*Func

	Depends []*File
	UsedBy  []*File
}
