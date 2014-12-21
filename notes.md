# TODO:

- add ring protection in page table entries

# Asm grammar

```
import (
    "pack"
)

const (
    Struct:field = 3
    Struct:
    Struct:function =
)

func Struct:main {

.begin
    addi r0 3
    mul r0 r1 r1
    div r0 r2 r3
    andi r0 0x3
    slt r0 r3 r4
    ld r0 r0 0
    lbu r3 r0 -3
    br r0 r0 .begin // really?
    j .end

.end
    j main

}
```

# Thoughts on asm8


