The EnableInt/DisableInt member functions of the Interrupt struct might be
redundant.  Programs running on CPU will just directly manipulate the memory.
And programs outside CPU do not need to modify that.

But let's just leave it there for a while. If they have no use, we will remove
them later.


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
