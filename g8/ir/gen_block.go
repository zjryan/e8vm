package ir

var basicOpFuncs = map[string]func(dest, r1, r2 uint32) uint32{
	"+": asm.add,
	"-": asm.sub,
	"*": asm.mul,
	"/": asm.div,
	"%": asm.mod,
	"&": asm.and,
	"|": asm.or,
}

func genOp(b *Block, op op) {
	switch op := op.(type) {
	case *arithOp:
		if op.dest == nil {
			panic("arith with no destination")
		}

		if op.a != nil {
			// binary arith op
			loadRef(b, _1, op.a)
			loadRef(b, _2, op.b)

			fn := basicOpFuncs[op.op]
			if fn != nil {
				b.inst(fn(_1, _1, _2))
			} else {
				panic("unknown arith op: " + op.op)
			}

			saveRef(b, _1, op.dest)
		} else {
			// unary arith op
			switch op.op {
			case "":
				loadRef(b, _1, op.b)
			case "-":
				loadRef(b, _1, op.b)
				b.inst(asm.sub(_1, _0, _1))
			default:
				panic("unkown arith unary op: " + op.op)
			}

			saveRef(b, _1, op.dest)
		}
	default:
		panic("unknown op type")
	}
}

func genJump(b *Block, j *jump) {
	panic("todo")
}

func genBlock(b *Block) {
	for _, op := range b.ops {
		genOp(b, op)
	}

	for _, jump := range b.jumps {
		genJump(b, jump)
	}
}
