package asm8

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func pf(s string) {
	r := strings.NewReader(s)
	rc := ioutil.NopCloser(r)
	p := newParser("t.s8", rc)
	p.ParseFunc = parseFunc
	var fs []*Func

	for {
		b := p.Block()
		if b == nil {
			break
		}
		f := b.(*Func)
		fs = append(fs, f)
	}

	errs := p.Errs()
	if errs != nil {
		for _, e := range errs {
			fmt.Println(e)
		}
	} else {
		for _, f := range fs {
			fmt.Printf("func %s {\n", f.name.Lit)
			for _, line := range f.Lines {
				for i, op := range line.Ops {
					if i == 0 {
						fmt.Print("    ")
					} else {
						fmt.Print(" ")
					}

					fmt.Print(op.Lit)
				}
				fmt.Println()
			}

			fmt.Printf("}\n")
		}
	}
}

func ExampleFunc_1() {
	pf(`
	func main {
		lui r4 /*inline comment*/ something

		// blank lines are ignored
		lui a5   anything		cool // some comment
		/* some block comment also */
	}`)
	// Output:
	// func main {
	//     lui r4 something
	//     lui a5 anything cool
	// }
}

func ExampleFunc_2() {
	pf(`
	func main {}
	`)
	// Output:
	// t.s8:2: expect end-line, got '}'
}

func ExampleFunc_3() {
	pf(`
	func main {
	}
	`)
	// Output:
	// func main {
	// }
}

func ExampleFunc_4() {
	pf(`
	func main t {
	}
	`)
	// Output:
	// t.s8:2: expect '{', got operand
	// t.s8:3: expect keyword func, got '}'
}

func ExampleFunc_5() {
	pf(`
	func main {
		a "something" "key"
		b
	}
	`)
	// Output:
	// t.s8:3: expect operand, got string
}
