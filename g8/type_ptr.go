package g8

type typPtr struct{ t typ } // a pointer type
type typSlice struct{ t typ }
type typArray struct {
	t typ
	n uint32
}
