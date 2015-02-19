var data {
    x 48 65 6c 6c 6f 0a 00
}

func main {
    xor r0 r0 r0
    lui r3 data
    ori r3 r3 data

.loop
    lb r1 r3
    beq r1 r0 .end
    jal fmt2.PrintChar
    addi r3 r3 1
    j .loop

.end
    halt
}
