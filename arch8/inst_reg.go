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
		case 0: // sll
			d = s1 << shift
		case 1: // srl
			d = s1 >> shift
		case 2: // sra
			d = uint32(int32(s1) >> shift)
		case 3: // sllv
			d = s1 << s2
		case 4: // srlv
			d = s1 >> s2
		case 5: // srla
			d = uint32(int32(s1) >> s2)
		case 6: // add
			d = s1 + s2
		case 7: // sub
			d = s1 - s2
		case 8: // and
			d = s1 & s2
		case 9: // or
			d = s1 | s2
		case 10: // xor
			d = s1 ^ s2
		case 11: // nor
			d = ^(s1 | s2)
		case 12: // slt
			if int32(s1) < int32(s2) {
				d = 1
			} else {
				d = 0
			}
		case 13: // sltu
			if s1 < s2 {
				d = 1
			} else {
				d = 0
			}
		case 14: // mul
			d = uint32(int32(s1) * int32(s2))
		case 15: // mulu
			d = s1 * s2
		case 16: // div
			if s2 == 0 {
				d = 0
			} else {
				d = uint32(int32(s1) / int32(s2))
			}
		case 17: // divu
			if s2 == 0 {
				d = 0
			} else {
				d = s1 / s2
			}
		case 18: // mod
			if s2 == 0 {
				d = 0
			} else {
				d = uint32(int32(s1) % int32(s2))
			}
		case 19: // modu
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
		case 0: // fadd
			fd = f1 + f2
		case 1: // fsub
			fd = f1 - f2
		case 2: // fmul
			fd = f1 * f2
		case 3: // fdiv
			fd = f1 / f2
		case 4: // fint
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
