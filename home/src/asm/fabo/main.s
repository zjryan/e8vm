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