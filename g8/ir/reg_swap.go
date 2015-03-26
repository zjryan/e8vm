package ir

func saveVar(b *Block, reg uint32, v *stackVar) {
	// TODO: save the the stack variable back to the frame
}

func loadVar(b *Block, reg uint32, v *stackVar) {

}

func saveRef(b *Block, reg uint32, r ref) {
	switch r := r.(type) {
	case *stackVar:
		saveVar(b, reg, r)
	default:
		panic("not implemented")
	}
}

func loadRef(b *Block, reg uint32, r ref) {
	switch r := r.(type) {
	case *stackVar:
		loadVar(b, reg, r)
	default:
		panic("not implemented")
	}
}
