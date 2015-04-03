package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"io/ioutil"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/g8"
	"lonnie.io/e8vm/lex8"
)

func exit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	os.Exit(-1)
}

func printErrs(es []*lex8.Error) {
	if len(es) == 0 {
		return
	}
	for _, e := range es {
		fmt.Println(e)
	}
	exit(nil)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		exit(errors.New("need exactly one input input file"))
	}

	fname := args[0]

	input, e := ioutil.ReadFile(fname)
	if e != nil {
		exit(e)
	}

	bs, es := g8.CompileBareFunc(fname, string(input))
	printErrs(es)

	ncycle, e := arch8.RunImage(bs, 100000)
	fmt.Printf("(%d cycles)\n", ncycle)
	if e != nil {
		fmt.Println(e)
	}
}
