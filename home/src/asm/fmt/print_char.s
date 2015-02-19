// print the char in r1
func PrintChar {
    addi    r2 r0 0x2000
.loop
    lbu     r4 r2 1
    bne     r4 r0 .loop // wait for invalid

    sb      r1 r2

    addi    r1 r0 1
    sb      r1 r2 1

    mov     pc ret
}