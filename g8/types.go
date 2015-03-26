package g8

type typ interface{}

type typErr struct{} // for storing error types

type typUint struct{} // uint32
type typInt struct{}  // int32, default type for all bare integers
type typUint8 struct{}
type typInt8 struct{}
type typFloat32 struct{}
type typStr struct{} // string constant

type typPtr struct{ t typ } // a pointer type
type typSlice struct{ t typ }
type typArray struct {
	t typ
	n uint32
}

type varDecl struct {
	name string
	t    typ
}

type funcDecl struct {
	name string
	t    *typFunc
}

type typStruct struct {
	fields  []*varDecl
	methods []*funcDecl
}

type typFunc struct {
	argTypes []typ
	regTypes []typ

	// optional names
	argNames []string
	retNames []string
}
