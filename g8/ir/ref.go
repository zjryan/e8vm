package ir

type ref interface{}

// stackVar is a variable on stack
type stackVar struct {
	name   string
	id     int
	offset int32
	size   int32

	// regOnly stack vars does not take frame space on the stack
	regOnly bool
}

type heapVar struct{ pkg, sym int } // a variable symbol on heap
type funcSym struct{ pkg, sym int } // a function symbol
type number struct{ v uint32 }      // a constant number
