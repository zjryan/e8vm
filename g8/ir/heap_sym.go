package ir

// a variable on heap
type heapSym struct {
	pkg, sym uint32
	size     uint32
}

func (s *heapSym) Addressable() bool { return false }
