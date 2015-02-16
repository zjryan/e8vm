var data {
    u32 0x6c6c6548 0x0a6f
}

func main {
    xor r0 r0 r0
    lui r3 data
    ori r3 r3 data

.loop
    lb r1 r3
    beq r1 r0 .end
    jal printChar
    addi r3 r3 1
    j .loop

.end
    halt
}

func printChar {
    addi r2 r0 0x2000
.loop
    lbu r4 r2 1
    bne r4 r0 .loop
    sb r1 r2
    addi r1 r0 1
    sb r1 r2 1
    mov pc ret
}

