package parse

import (
	"testing"

	"strings"
)

func TestFile_good(t *testing.T) {
	for _, s := range []string{
		"func f() {}",
		"func f() {\n}",
		"func f(int) {}",
		"func f(a int) {}",
		"func f(a int,) {}",
		"func f(a, b int) {}",
		"func f(a int, b int) {}",
		"func f() (int) {}",
		"func f() (a int) {}",
		"func f() (a, b int) {}",
		"func f() (a int, b int) {}",
		"func f(int) (a int, b int) {}",
		"func f(int) (a int, b int,) {}",
		`func f(int) (
			a int, 
			b int,
		) {}
		`,
	} {
		buf := strings.NewReader(s)
		f, es := File("test.g", buf)
		if es != nil {
			t.Log(s)
			for _, e := range es {
				t.Log(e)
			}
			t.Fail()
		} else if f == nil {
			t.Log(s)
			t.Log("returned nil")
			t.Fail()
		}
	}
}

func TestFile_bad(t *testing.T) {
	for _, s := range []string{
		"func f() ",
		"func f {}",
		"func f(",
		"func f)",
		"func f; {}",
		"func f(a int) () {}",
		"func f(,) {}",
		"func f(,a) {}",
		"func f(a int) (,a) {}",
		"func f(a int \n b int) {}",
		"func f() \n {}",
	} {
		buf := strings.NewReader(s)
		_, es := File("test.g", buf)
		if es == nil {
			t.Log(s)
			t.Log("should fail")
			t.Fail()
		}
	}
}
