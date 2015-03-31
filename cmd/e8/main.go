package main

import (
	"fmt"
	"os"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/g8"
)

func main() {
	b := build8.NewBuilder(".")
	b.Verbose = true
	b.AddLang("asm", asm8.Lang())
	b.AddLang("bare", g8.BareFunc())
	b.AddLang("", g8.Lang())

	es := build8.BuildAll(b)
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		os.Exit(-1)
	}
}
