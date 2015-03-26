import {
    "asm/fmt"
}

var data {
    u32 0x6c6c6548 0x0a6f
}

func main {
    xor     r0 r0 r0 // clear r0
    lui     sp 0x1000
    addi    sp sp 4096 // set sp the stack

    lui r3 data
    ori r3 r3 data

.loop
    lb r1 r3
    beq r1 r0 .end
    jal fmt.PrintChar
    addi r3 r3 1
    j .loop

.end
    halt
}
