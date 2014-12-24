# TODO:

- add ring protection in page table entries

# Asm grammar

```
import {
    "pack"
}

const {
    StructName:field 0
    StructName:dfadf 0
}

var somedata {
    align 4
    x 33 32 32 43
    x 32 33 44 32
    i32 32 32
    i32 32 -32 0x32 
    i8 32 33
    "something" 0
}

func StructName:Hello {

.begin
    addi r0 3
    mul r0 r1 r1
    div r0 r2 r3
    andi r0 0x3
    slt r0 r3 r4
    ld r0 r0 0
    lbu r3 r0 -3
    beq r0 r0 .begin
    j .end

.end
    j main
}
```

# Thoughts on asm8


