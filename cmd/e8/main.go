package main

import (
	"fmt"
	"os"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/build8"
	"lonnie.io/e8vm/g8"
)

func main() {
	home := build8.NewFileHome(".", g8.Lang())
	home.AddLang("asm", asm8.Lang())
	home.AddLang("bare", g8.BareFunc())

	b := build8.NewBuilder(home)
	b.Verbose = true

	es := b.BuildAll()
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		os.Exit(-1)
	}
}
