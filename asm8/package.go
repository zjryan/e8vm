package asm8

type Package struct {
	name string
	path string

	imports  []*Package
	symNames []string
	symObjs  []interface{}

	symTable *SymTable
	symIndex map[string]uint32
	pkgIndex map[string]uint32

	funcs  []*funcObj
	vars   []*varObj
	consts []*constObj
}

type varObj struct{}
type constObj struct{}
