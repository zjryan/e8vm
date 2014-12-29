package link8

type Symbol struct {
	Name string
	Type int
}

const (
	SymFunc = iota
	SymVar
)
