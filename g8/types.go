package g8

type typPtr struct{ t typ } // a pointer type
type typSlice struct{ t typ }
type typArray struct {
	t typ
	n uint32
}

type varSig struct {
	name string
	t    typ
}

type funcSig struct {
	name string
	t    *typFunc
}

type typStruct struct {
	fields  []*varSig
	methods []*funcSig
}

type typFunc struct {
	argTypes []typ
	retTypes []typ

	// optional names
	argNames []string
	retNames []string
}
