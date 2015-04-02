package ir

import (
	"fmt"
)

// Block is a basic block
type Block struct {
	id    int // basic block ida
	ops   []op
	jumps []*jump

	insts   []*inst
	spMoved bool

	frameSize *int32

	next *Block // next in the link
}

func (b *Block) String() string { return fmt.Sprintf("B%d:", b.id) }

func (b *Block) addOp(op op) { b.ops = append(b.ops, op) }

// SetNext sets the next block to run after this block
func (b *Block) SetNext(next *Block) { b.next = next }

// Arith append an arithmetic operation to the basic block
func (b *Block) Arith(dest Ref, x Ref, op string, y Ref) {
	b.addOp(&arithOp{dest, x, op, y})
}

// Assign appends an assignment operation to the basic block
func (b *Block) Assign(dest Ref, a Ref) {
	b.Arith(dest, nil, "", a)
}

// Call appends a function call operation to the basic block
func (b *Block) Call(dests []Ref, f Ref, sig *FuncSig, args ...Ref) {
	b.addOp(&callOp{dests, f, sig, args})
}

func (b *Block) addJump(j *jump) { b.jumps = append(b.jumps, j) }

// Jump appends a redirection to the end of the basic block.
// The redirection points to dest.
func (b *Block) Jump(dest *Block, x Ref, op string, y Ref) {
	b.addJump(&jump{x, op, y, dest.id})
}

// JumpID appends a redirection to the end of the basic block.
// The redirection points to the basic block of the particular id.
func (b *Block) JumpID(id int, x Ref, op string, y Ref) {
	b.addJump(&jump{x, op, y, id})
}

func (b *Block) inst(i uint32) *inst {
	ret := new(inst)
	ret.inst = i
	b.insts = append(b.insts, ret)
	return ret
}
