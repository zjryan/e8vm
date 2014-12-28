package asm8

type Package struct {
	name string
	path string

	imports map[uint32]string
	symbols map[uint32]interface{}
	funcs   []*funcObj
	vars    []*varObj
	consts  []*constObj
}

type funcObj struct{}
type varObj struct{}
type constObj struct{}
