package ir

// Block is a basic block
type Block struct {
	id    int // basic block ida
	ops   []op
	jumps []*jump

	insts []*inst
}

func (b *Block) addOp(op op) { b.ops = append(b.ops, op) }

func (b *Block) arith(dest ref, x ref, op string, y ref) {
	b.addOp(&arithOp{dest, x, op, y})
}

func (b *Block) assign(dest ref, a ref) {
	b.arith(dest, a, "", nil)
}

func (b *Block) call(dest ref, f ref, args ...ref) {
	b.addOp(&callOp{dest, f, args})
}

func (b *Block) addJump(j *jump) { b.jumps = append(b.jumps, j) }

func (b *Block) jump(dest *Block, x ref, op string, y ref) {
	b.addJump(&jump{x, op, y, dest.id})
}

func (b *Block) jumpID(id int, x ref, op string, y ref) {
	b.addJump(&jump{x, op, y, id})
}

func (b *Block) inst(i uint32) *inst {
	ret := new(inst)
	ret.inst = i
	b.insts = append(b.insts, ret)
	return ret
}
