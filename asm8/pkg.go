package asm8

type pkg struct {
	path  string
	files []*file

	imports *importDecl
}
