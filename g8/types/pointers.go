package types

import (
	"fmt"
)

// Pointer is a pointer type
type Pointer struct{ T T } // a pointer type

// String returns "*T"
func (t *Pointer) String() string { return "*" + t.T.String() }

// Size returns the address length of the architecture.
func (t *Pointer) Size() int32 { return 4 }

// Slice is a slice type
type Slice struct{ T T }

// String returns "[]T"
func (t *Slice) String() string { return "[]" + t.T.String() }

// Size returns the size of the slice, which is same as a pointer,
// because slices are reference.
func (t *Slice) Size() int32 { return 4 }

// Array is an array type of fixed size
type Array struct {
	T T
	N int32
}

// Size returns the total size of the array.
func (t *Array) Size() int32 { return t.T.Size() * t.N }

// String returns "[N]T"
func (t *Array) String() string {
	return fmt.Sprintf("[%d]%s", t.N, t.T)
}
