package arch8

// imm instructions
const (
	ADDI = 1
	SLTI = 2
	ANDI = 3
	ORI  = 4
	LUI  = 5

	LW  = 6
	LB  = 7
	LBU = 8
	SW  = 9
	SB  = 10
)

// reg instructions
const (
	SLL  = 0
	SRL  = 1
	SRA  = 2
	SLLV = 3
	SRLV = 4
	SRLA = 5
	ADD  = 6
	SUB  = 7
	AND  = 8
	OR   = 9
	XOR  = 10
	NOR  = 11
	SLT  = 12
	SLTU = 13
	MUL  = 14
	MULU = 15
	DIV  = 16
	DIVU = 17
	MOD  = 18
	MODU = 19

	FADD = 0
	FSUB = 1
	FMUL = 2
	FDIV = 3
	FINT = 4
)

// branch instructions
const (
	BNE = 32
	BEQ = 33
)

// system instructions
const (
	HALT    = 64
	SYSCALL = 65
	USERMOD = 66
	VTABLE  = 67
	IRET    = 68
	CPUID   = 69
)
