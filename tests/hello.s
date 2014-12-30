var msg {
    str     "Hello, world!\n\x00"
}

func main {
    xor     r0 r0 r0 // clear r0
    lui     r3 msg
    ori    r3 r3 msg

.loop
    lb      r1 r3
    beq     r1 r0 .end
    jal     printChar
    addi    r3 r3 1 // inc
    j       .loop

.end
    halt
}

// print the char in r1
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

