package main

import (
	"bytes"
	"flag"
	"fmt"
	// "io/ioutil"
	// "strings"
	"os"

	"lonnie.io/e8vm/arch8"
	"lonnie.io/e8vm/asm8"
	"lonnie.io/e8vm/dasm8"
	// "lonnie.io/e8vm/lex8"
)

var (
	doDasm      = flag.Bool("d", false, "do dump")
	ncycle      = flag.Int("n", 100000, "max cycles to execute")
	memSize     = flag.Int("m", 1<<30, "memory size")
	printStatus = flag.Bool("s", false, "print status after execution")
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
	if *printStatus {
		m.PrintCoreStatus()
	}

	if exp == nil {
		return ret, nil
	}
	return ret, exp
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "need exactly one input file\n")
		os.Exit(-1)
	}

	f, e := os.Open(args[0])
	if e != nil {
		fmt.Fprintf(os.Stderr, "open: %s", e)
		os.Exit(-1)
	}

	bs, es := asm8.BuildBareFunc(args[0], f)
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
