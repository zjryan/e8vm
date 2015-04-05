package parse

import (
	"testing"

	"strings"

	"lonnie.io/e8vm/g8/ast"
)

func TestStmts_good(t *testing.T) {
	// should be expressions
	for _, s := range []string{
		"_",
		"3",
		"-3",
		"-a",
		"!a",
		"0",
		"_a",
		"a",
		"a + 3",
		"print(3)",
		"print(3, 4)",
		"print()",
		"a",
		"a+3+4",
		"a * 3",
		"(a)",
		"((a))",
		"(a+3)*4",
		"4 * (a + 3)",
		"a == 4",
		"a > 5",
		"a < 6",
		"a >= 5",
		"a <= 6",
		"a != 7",
		"(void)()",
	} {
		buf := strings.NewReader(s)
		stmts, es := Stmts("test.g", buf)
		if es != nil {
			t.Log(s)
			for _, e := range es {
				t.Log(e)
			}
			t.Fail()
		} else if len(stmts) != 1 {
			t.Log(s)
			t.Error("should be one statement")
		} else {
			s := stmts[0]
			if _, ok := s.(*ast.ExprStmt); !ok {
				t.Log(s)
				t.Error("should be an expression")
			}
		}
	}

	// should be a statement
	for _, s := range []string{
		";",
		"{;;;;}",
		"{}",
		"{};",
		"{;}",
		"{3}",
		"a = 3",
		"a, b = 4, x",
		"a(), b() = 4, x(x())",
		"a := 3",
		"a := 3+4",
		"a, b := 4, x",
		"for true { }",
		"for (true) { }",
		"for a == 3 { }",
		"if true { }",
		"if (true) { }",
		"if true { } else { }",
		`if true {
			print(3)
			print(5)
		} else {
			print(4)
		}`,
		`for true {
			print(3)
			read()
		}`,
		"var a int",
		"var a int = 3",
		"var a = 3",
		"var a, b int = 3",
		"var a, b int",
		"var a, b int = 3, 4",
		"var ()",
		"var (a, b int)",
		"var (a, b int = 3, 4)",
		"var (a, b = 3, 4)",
		"var (a int; b int)",
		"var (a int\n b int)",
		"var (\n a int \n);",
	} {
		buf := strings.NewReader(s)
		stmts, es := Stmts("test.g", buf)
		if es != nil {
			t.Log(s)
			for _, e := range es {
				t.Log(e)
			}
			t.Fail()
		} else if len(stmts) != 1 {
			t.Log(s)
			t.Error("should be one statement")
		} else {
			s := stmts[0]
			if _, ok := s.(*ast.ExprStmt); ok {
				t.Log(s)
				t.Error("should not be an expression")
			}
		}
	}
}

func TestStmts_bad(t *testing.T) {
	// should be broken
	for _, s := range []string{
		"3 3",
		"3a",
		"3x3",
		"p(",
		"p(;)",
		"p())",
		"{",
		"}",
		"{}}",
		"()",
		"if true { ",
		"if true; { }",
		"if { }",
		"if true { else }",
		"if true { }; else {}",
		"if true else {}",
		"if true {} else; { }",
		"for ; {}",
		"for ; ",
		"for true ;",
		"if true { x( } else {}",
		"if true { x{ } else {}",
		"if true { { } else {}",
		"if true { x(;) } else {}",
		"var a",
		"var = 3",
		"var a b c",
		"var a b c = 3, 4",
		"var a b = c d",
		"var \n ()",
		"var (a = 3, b = 4)",
		"var (a)",
		"var {}",
		"var a)",
		"var ( a int;",
	} {
		buf := strings.NewReader(s)
		stmts, es := Stmts("test.g", buf)
		if es == nil || stmts != nil {
			t.Log(s)
			t.Error("should fail")
		}
	}
}
