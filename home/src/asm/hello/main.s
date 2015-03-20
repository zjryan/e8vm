var msg {
    str     "Hello, world!\n\x00"
}

func main {
    xor     r0 r0 r0 // clear r0
    lui     r3 msg
    ori     r3 r3 msg

.loop
    lb      r1 r3
    beq     r1 r0 .end
    jal     fmt.PrintChar
    addi    r3 r3 1 // inc
    j       .loop

.end
    halt
}
