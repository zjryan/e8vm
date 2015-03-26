package arch8

import (
	"math"
)

// InstReg executes register based instructions
type instReg struct{}

// I executes the register instruction.
// Might return invalid instruction exception.
func (i *instReg) I(cpu *cpu, in uint32) *Excep {
	// bit (32:24] == 0
	if ((in >> 24) & 0xff) != 0 {
		panic("not a register inst")
	}

	dest := (in >> 21) & 0x7   // (24:21]
	src1 := (in >> 18) & 0x7   // (21:18]
	src2 := (in >> 15) & 0x7   // (18:15]
	shift := (in >> 10) & 0x1f // (15:10]
	isFloat := (in >> 8) & 0x1 // (9:8]
	funct := in & 0xff         // (8:0]

	s1 := cpu.regs[src1]
	s2 := cpu.regs[src2]
	d := uint32(0)

	if isFloat == 0 {
		switch funct {
		case SLL:
			d = s1 << shift
		case SRL:
			d = s1 >> shift
		case SRA:
			d = uint32(int32(s1) >> shift)
		case SLLV:
			d = s1 << s2
		case SRLV:
			d = s1 >> s2
		case SRLA:
			d = uint32(int32(s1) >> s2)
		case ADD:
			d = s1 + s2
		case SUB:
			d = s1 - s2
		case AND:
			d = s1 & s2
		case OR:
			d = s1 | s2
		case XOR:
			d = s1 ^ s2
		case NOR:
			d = ^(s1 | s2)
		case SLT:
			if int32(s1) < int32(s2) {
				d = 1
			} else {
				d = 0
			}
		case SLTU:
			if s1 < s2 {
				d = 1
			} else {
				d = 0
			}
		case MUL:
			d = uint32(int32(s1) * int32(s2))
		case MULU:
			d = s1 * s2
		case DIV:
			if s2 == 0 {
				d = 0
			} else {
				d = uint32(int32(s1) / int32(s2))
			}
		case DIVU:
			if s2 == 0 {
				d = 0
			} else {
				d = s1 / s2
			}
		case MOD:
			if s2 == 0 {
				d = 0
			} else {
				d = uint32(int32(s1) % int32(s2))
			}
		case MODU:
			if s2 == 0 {
				d = 0
			} else {
				d = s1 % s2
			}
		default:
			return errInvalidInst
		}
	} else {
		f1 := math.Float32frombits(s1)
		f2 := math.Float32frombits(s2)
		var fd float32
		switch funct {
		case FADD:
			fd = f1 + f2
		case FSUB:
			fd = f1 - f2
		case FMUL:
			fd = f1 * f2
		case FDIV:
			fd = f1 / f2
		case FINT:
			d = uint32(f1)
		default:
			return errInvalidInst
		}

		if funct != 4 {
			d = math.Float32bits(fd)
		}
	}

	cpu.regs[dest] = d
	return nil
}
