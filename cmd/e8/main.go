package main

import (
	"fmt"
	"os"

	"lonnie.io/e8vm/build8"
)

func main() {
	es := build8.BuildAll(".")
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		os.Exit(-1)
	}
}
