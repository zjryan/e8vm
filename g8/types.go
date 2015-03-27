package g8

type typ interface{}

type typBasic int

const (
	typErr typBasic = iota
	typInt
	typUint
	typInt8
	typUint8
	typFloat32
	typString
)

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
	regTypes []typ

	// optional names
	argNames []string
	retNames []string
}
