// a char is sent in via r1
func PrintChar {
	// use r2 and r3
	addi sp sp -8
	sw r2 sp
	sw r3 sp 4

	ori r2 r0 0x2000 // the address of serial port
.wait
	lbu r3 r2 1
	bne r3 r0 .wait // wait for invalid

	sb r1 r2
	ori r3 r0 1 // set r3 to 1
	sb r3 r2 1

	// restore r2 and r3
	lw r2 sp
	lw r3 sp 4

	addi sp sp 8
	mov pc ret
}

// Print a 32-bit unsigned integer
func PrintUint32 {
	// saving used registers
	sw ret sp -4
	addi sp sp -24
	sw r1 sp
	sw r2 sp 4
	sw r3 sp 8
	sw r4 sp 12
	
	bne r1 r0 .nonzero
.zero
	addi r1 r0 0x30 // '0'
	jal PrintChar
	j .end

.nonzero
	ori r2 r0 10 // r2 = 10
	ori r3 r0 1 // r3 = 1, will be used as the divisor
.find
	mulu r3 r3 r2
	slt r4 r1 r3
	beq r4 r0 .find

	divu r3 r3 r2
.loop
	sw r1 sp 16 // save r1
	divu r1 r1 r3 // got the digit
	addi r1 r1 0x30 // convert digit to char

	jal PrintChar

	lw r1 sp 16 // load back r1
	modu r1 r1 r3
	div r3 r3 r2 // next order of magnitude
	bne r3 r0 .loop

.end
	addi r1 r0 0xa
	jal PrintChar

	lw r1 sp 
	lw r2 sp 4
	lw r3 sp 8
	lw r4 sp 12
	addi sp sp 24
	lw pc sp -4
}

// TODO: print the sign number
// Print a 32-bit signed integer
func PrintInt32 {
	// saving used registers
	sw ret sp -4
	addi sp sp -24
	sw r1 sp
	sw r2 sp 4
	sw r3 sp 8
	sw r4 sp 12
	
	bne r1 r0 .nonzero
.zero
	addi r1 r0 0x30 // '0'
	jal PrintChar
	j .end

.nonzero
	ori r2 r0 10 // r2 = 10
	ori r3 r0 1 // r3 = 1, will be used as the divisor
.find
	mulu r3 r3 r2
	slt r4 r1 r3
	beq r4 r0 .find

	divu r3 r3 r2
.loop
	sw r1 sp 16 // save r1
	divu r1 r1 r3 // got the digit
	addi r1 r1 0x30 // convert digit to char

	jal PrintChar

	lw r1 sp 16 // load back r1
	modu r1 r1 r3
	div r3 r3 r2 // next order of magnitude
	bne r3 r0 .loop

.end
	addi r1 r0 0xa
	jal PrintChar

	lw r1 sp 
	lw r2 sp 4
	lw r3 sp 8
	lw r4 sp 12
	addi sp sp 24
	lw pc sp -4
}
