func main {
    xor     r0 r0 r0 // clear r0

    addi    r1 r0 15
    lui     sp 0x1000
    addi    sp sp 4096 // set sp
    sw      r1 sp
    addi    sp sp 8
    
    jal     fabo
    lw      r1 sp -4
    addi    sp sp -8
    
    sw      r1 sp
    addi    sp sp 4
    jal     printNum
    addi    sp sp -4
    
    lw      r1 sp

    halt
}

// fabonacci func
func fabo {
    sw      ret sp
    lw      r1 sp -8
    beq     r1 r0 .ret0
    addi    r1 r1 -1
    beq     r1 r0 .ret1

    sw      r1 sp 8  // arg for recursive call
    addi    sp sp 16
    jal     fabo
    lw      r2 sp -4
    addi    sp sp -16
    sw      r2 sp 4  // save the return value

    lw      r1 sp -8 // load the arg again
    addi    r1 r1 -2 // -2

    sw      r1 sp 8
    addi    sp sp 16
    jal     fabo
    lw      r2 sp -4
    addi    sp sp -16

    lw      r1 sp 4
    add     r1 r1 r2
    j       .out

.ret0
    mov     r1 r0
    j       .out

.ret1
    addi    r1 r0 1

.out
    sw      r1 sp -4
    lw      pc sp // return
}

func printNum {
    sw      ret sp
    lw      r1 sp -4
    bne     r1 r0 .nonzero
.zero
    addi    r1 r0 0x30 // '0'
    jal     printChar
    j       .end

.nonzero
    addi    r2 r0 10
    addi    r3 r0 1
.find
    mulu    r3 r3 r2
    slt     r4 r1 r3
    beq     r4 r0 .find

    divu    r3 r3 r2
.loop
    sw      r1 sp 4 // save r1
    divu    r1 r1 r3
    addi    r1 r1 0x30 // convert digit to char
    
    sw      r2 sp 8
    jal     printChar
    lw      r2 sp 8

    lw      r1 sp 4 // load back r1
    modu    r1 r1 r3
    divu    r3 r3 r2
    bne     r3 r0 .loop

.end
    addi    r1 r0 0xa
    jal     printChar

    lw      pc sp // return
}

func printChar {
    addi    r2 r0 0x2000
.loop
    lbu     r4 r2 1
    bne     r4 r0 .loop // wait for invalid

    sb      r1 r2
    
    addi    r1 r0 1
    sb      r1 r2 1
    
    mov     pc ret
}
