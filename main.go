package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/dasm8"
	"lonnie.io/e8vm/lex8"
)

func buildBareFunc(file, code string) ([]byte, []*lex8.Error) {
	rc := ioutil.NopCloser(strings.NewReader(code))
	return asm8.BuildBareFunc(file, rc)
}

var code = `

.main
	xor 	r0 r0 r0 // clear r0

	addi 	r1 r0 10
	lui     sp 0x1000
	addi 	sp sp 4096 // set sp
	sw		r1 sp 0
	addi  	sp sp 8

	jal		.fabo
	lw 		r1 sp -4
	addi	sp sp -8

	halt

.fabo
	sw		ret sp 0
	lw		r1 sp -8
	beq		r1 r0 .ret0
	addi    r1 r1 -1
	beq		r1 r0 .ret1

	sw		r1 sp 8  // arg for recursive call
	addi	sp sp 16
	jal		.fabo
	lw		r2 sp -4
	addi	sp sp -16
	sw		r2 sp 4  // save the return value

	lw		r1 sp -8 // load the arg again
	addi	r1 r1 -2 // -2

	sw		r1 sp 8 
	addi 	sp sp 16
	jal		.fabo
	lw		r2 sp -4
	addi    sp sp -16

	lw		r1 sp 4
	add		r1 r1 r2
	j 		.out

.ret0
	mov     r1 r0
	j       .out

.ret1
	addi 	r1 r0 1

.out
	sw		r1 sp -4
	lw		pc sp 0 // return

`

var (
	doDasm  = flag.Bool("d", false, "do dump")
	ncycle  = flag.Int("n", 100000, "max cycles to execute")
	memSize = flag.Int("m", 1<<30, "memory size")
)

func run(bs []byte) (int, error) {
	r := bytes.NewBuffer(bs)

	// create a single core machine
	m := arch8.NewMachine(uint32(*memSize), 1)
	e := m.LoadImage(r, arch8.InitPC)
	if e != nil {
		return 0, e
	}

	ret, exp := m.Run(*ncycle)
	m.PrintCoreStatus()

	if exp == nil {
		return ret, nil
	}
	return ret, exp
}

func main() {
	bs, es := buildBareFunc("test.s", code)
	if len(es) > 0 {
		for _, e := range es {
			fmt.Println(e)
		}
	} else {
		if *doDasm {
			lines := dasm8.Dasm(bs, arch8.InitPC)
			for _, line := range lines {
				fmt.Println(line)
			}
		} else {
			n, e := run(bs)
			fmt.Printf("(%d cycles)\n", n)
			if e != nil {
				fmt.Println(e)
			}
		}
	}
}
