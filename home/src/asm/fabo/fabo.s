import {
	"asm/fmt"
}

func main {
    xor     r0 r0 r0 // clear r0

    lui     sp 0x1000
    addi    sp sp 4096 // set sp the stack

    addi    r1 r0 17
	jal     fabo 
    // r1 = fabo(n) now
    jal     fmt.PrintUint32

    halt
}

// fabonacci func
func fabo {
    beq     r1 r0 .ret0

    addi    r1 r1 -1
    beq     r1 r0 .ret1
    
    sw      ret sp -4
    addi    sp sp -12
    sw      r1 sp // n-1 saved here
    sw      r2 sp 4

    jal     fabo

    mov     r2 r1 // f(n-1)

    lw      r1 sp
    addi    r1 r1 -1 // r1 = n-2 now
    jal     fabo
    add     r1 r1 r2
    
    lw      r2 sp 4 // restore r2
    addi    sp sp 12
    lw      pc sp -4 // load return

.ret1
    addi    r1 r0 1
    mov     pc ret

.ret0
    mov     r1 r0
    mov     pc ret 
}
