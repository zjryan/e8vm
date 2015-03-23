package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"lonnie.io/e8vm/g8/parse"
)

func exit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	os.Exit(-1)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		exit(errors.New("need exactly one input input file"))
	}

	fname := args[0]

	fin, e := os.Open(fname)
	if e != nil {
		exit(e)
	}

	toks, es := parse.Tokens(fname, fin)
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		exit(nil)
	}

	for _, t := range toks {
		fmt.Println(t)
	}
}
