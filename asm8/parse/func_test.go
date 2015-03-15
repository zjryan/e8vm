package parse

import (
	"fmt"
	"io/ioutil"
	"strings"

	"lonnie.io/e8vm/asm8/ast"
	"lonnie.io/e8vm/lex8"
)

func pf(s string) {
	r := strings.NewReader(s)
	rc := ioutil.NopCloser(r)
	p := newParser("t.s8", rc)
	var fs []*ast.FuncDecl

	for {
		if p.see(lex8.EOF) {
			break
		}

		f := parseFunc(p)
		if f == nil {
			break
		}

		fs = append(fs, f)
	}

	errs := p.Errs()
	if errs != nil {
		for _, e := range errs {
			fmt.Println(e)
		}
	} else {
		for _, f := range fs {
			fmt.Printf("func %s {\n", f.Name.Lit)
			for _, stmt := range f.Stmts {
				for i, op := range stmt.Ops {
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
		add r4 /*inline comment*/ r3 r5

		// blank lines are ignored
		sub r0   r0		r1 // some comment
		/* some block comment also */
	}`)
	// Output:
	// func main {
	//     add r4 r3 r5
	//     sub r0 r0 r1
	// }
}

func ExampleFunc_2() {
	pf(`
	func main {}
	`)
	// Output:
	// func main {
	// }
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
}

func ExampleFunc_5() {
	pf(`
	func main {
		a "something" "key"
		j .lab 
	}
	`)
	// Output:
	// t.s8:3: expect operand, got string
}

func ExampleFunc_6() {
	pf(`func main { invalid }`)
	// Output:
	// t.s8:1: invalid asm instruction "invalid"
}

func ExampleFunc_7() {
	pf(`func main { j .lab:inv }`)
	// Output:
	// t.s8:1: invalid label: ".lab:inv"
}
