package ast

// File represents a file.
type File struct {
	Funcs []*FuncDecl
	Vars  []*VarDecl
}
