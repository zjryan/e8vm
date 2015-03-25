package ir

// Block is a basic block
type Block struct {
	id    int // basic block ida
	ops   []op
	jumps []*jump
}
