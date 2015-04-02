package ir

import (
	"fmt"

	"lonnie.io/e8vm/link8"
)

var basicOpFuncs = map[string]func(
	dest, r1, r2 uint32,
) uint32{
	"+":   asm.add,
	"-":   asm.sub,
	"*":   asm.mul,
	"/":   asm.div,
	"%":   asm.mod,
	"&":   asm.and,
	"|":   asm.or,
	"xor": asm.xor,
	"nor": asm.nor,
}

func genArithOp(b *Block, op *arithOp) {
	if op.dest == nil {
		panic("arith with no destination")
	}

	if op.a != nil {
		// binary arith op
		loadRef(b, _4, op.a)
		loadRef(b, _1, op.b)

		fn := basicOpFuncs[op.op]
		if fn != nil {
			b.inst(fn(_4, _4, _1))
		} else {
			switch op.op {
			case "==":
				b.inst(asm.xor(_4, _4, _1))  // the diff
				b.inst(asm.sltu(_4, _0, _4)) // if _4 is 0, _4 <= 0
				b.inst(asm.xori(_4, _4, 1))  // flip
			case "!=":
				b.inst(asm.xor(_4, _4, _1))  // the diff
				b.inst(asm.sltu(_4, _0, _4)) // if _4 is 0, _4 <= 0
			case ">":
				b.inst(asm.slt(_4, _1, _4))
			case "<":
				b.inst(asm.slt(_4, _4, _1)) // delta = b - a
			case ">=":
				b.inst(asm.slt(_4, _4, _1))
				b.inst(asm.xori(_4, _4, 1)) // flip
			case "<=":
				b.inst(asm.slt(_4, _1, _4))
				b.inst(asm.xori(_4, _4, 1)) // flip
			case "u>":
				b.inst(asm.sltu(_4, _1, _4))
			case "u<":
				b.inst(asm.sltu(_4, _4, _1))
			case "u>=":
				b.inst(asm.sltu(_4, _4, _1))
				b.inst(asm.xori(_4, _4, 1))
			case "u<=":
				b.inst(asm.sltu(_4, _1, _4))
				b.inst(asm.xori(_4, _4, 1))
			default:
				panic("unknown arith op: " + op.op)
			}
		}

		saveRef(b, _4, op.dest)
	} else {
		// unary arith op
		switch op.op {
		case "":
			loadRef(b, _4, op.b)
		case "-":
			loadRef(b, _4, op.b)
			b.inst(asm.sub(_4, _0, _4))
		case "!":
			b.inst(asm.sltu(_4, _0, _4)) // test non-zero first
			b.inst(asm.xori(_4, _4, 1))  // and flip
		case "?": // test if it is non-zero
			b.inst(asm.sltu(_4, _0, _4))
		default:
			panic("unkown arith unary op: " + op.op)
		}

		saveRef(b, _4, op.dest)
	}
}

func genCallOp(b *Block, op *callOp) {
	sig := op.sig

	// load the args
	for i, arg := range sig.args {
		if arg.viaReg == 0 {
			loadRef(b, _4, op.args[i]) // load the arg to r1
			saveArg(b, _4, arg)        // push it on the stack for calling
		}
	}
	for i, arg := range sig.args {
		if arg.viaReg > 0 {
			loadRef(b, arg.viaReg, op.args[i])
		}
	}

	// do the function call
	if s, ok := op.f.(*funcSym); ok {
		jal := b.inst(asm.jal(0))
		jal.sym = &linkSym{link8.FillLink, s.pkg, s.sym}
	} else {
		panic(fmt.Errorf("todo: calling function pointer: %T", op.f))
	}

	// unload the returns
	for i, ret := range sig.rets {
		if ret.viaReg > 0 {
			saveRef(b, ret.viaReg, op.dest[i])
		}
	}
	for i, ret := range sig.rets {
		if ret.viaReg == 0 {
			loadArg(b, _4, ret)
			saveRef(b, _4, op.dest[i])
		}
	}
}

func genOp(b *Block, op op) {
	switch op := op.(type) {
	case *arithOp:
		genArithOp(b, op)
	case *callOp:
		genCallOp(b, op)
	default:
		panic("unknown op type")
	}
}

func genBlock(b *Block) {
	for _, op := range b.ops {
		genOp(b, op)
	}

	if b.jump == nil {
		/* do nothing */
	} else if b.jump.typ == jmpAlways {
		b.jumpInst = b.inst(asm.j(0))
	} else if b.jump.typ == jmpIfNot {
		loadRef(b, _4, b.jump.cond)
		b.jumpInst = b.inst(asm.beq(_0, _4, 0))
	} else if b.jump.typ == jmpIf {
		loadRef(b, _4, b.jump.cond)
		b.jumpInst = b.inst(asm.bne(_0, _4, 0))
	}
}
