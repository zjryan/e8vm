package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"lonnie.io/e8vm/g8/parse"
	"lonnie.io/e8vm/lex8"
)

func exit(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
	}
	os.Exit(-1)
}

func printErrs(es []*lex8.Error) {
	if es != nil {
		for _, e := range es {
			fmt.Println(e)
		}
		exit(nil)
	}
}

func main() {
	doTokens := flag.Bool("toks", false, "parse tokens")
	doExpr := flag.Bool("expr", true, "parse as an expression")
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

	if *doTokens {
		toks, es := parse.Tokens(fname, fin)
		printErrs(es)
		for _, t := range toks {
			fmt.Printf("%s: %s %q\n", t.Pos, parse.TypeStr(t.Type), t.Lit)
		}
	} else if *doExpr {
		expr, es := parse.Expr(fname, fin)
		printErrs(es)
		fmt.Println(parse.PrintExpr(expr))
	}
}
