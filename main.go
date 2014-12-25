package main

import (
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
	mov r0 r1
	.lab
	j .l
`

func main() {
	bs, es := buildBareFunc("test.s", code)
	if len(es) > 0 {
		for _, e := range es {
			fmt.Println(e)
		}
	} else {
		lines := dasm8.Dasm(bs, arch8.InitPC)
		for _, line := range lines {
			fmt.Println(line)
		}
	}
}
