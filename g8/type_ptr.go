package g8

import (
	"fmt"
)

type typPtr struct{ t typ } // a pointer type

func (t *typPtr) String() string { return "*" + t.t.String() }
func (t *typPtr) Size() int32    { return 4 }

type typSlice struct{ t typ }

func (t *typSlice) String() string { return "[]" + t.t.String() }
func (t *typSlice) Size() int32    { panic("todo") }

type typArray struct {
	t typ
	n int32
}

func (t *typArray) Size() int32 { return 8 }
func (t *typArray) String() string {
	return fmt.Sprintf("[%d]%s", t.n, t.t)
}
