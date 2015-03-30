package main

import (
	"fmt"
	"os"

	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/build8"
)

func main() {
	es := build8.BuildAll(".", true, asm8.Lang())
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		os.Exit(-1)
	}
}
