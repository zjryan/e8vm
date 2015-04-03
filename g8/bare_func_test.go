package g8

import (
	"testing"

	"strings"

	"lonnie.io/e8vm/arch8"
)

func TestBareFunc_good(t *testing.T) {
	const N = 100000

	o := func(input, output string) {
		bs, es := CompileBareFunc("main.g", input)
		if es != nil {
			t.Log(input)
			for _, e := range es {
				t.Log(e)
			}
			t.Error("should not compile")
			return
		}

		ncycle, out, e := arch8.RunImageOutput(bs, N)
		if ncycle == N {
			t.Log(input)
			t.Error("running out of time")
			return
		}
		if !arch8.IsHalt(e) {
			t.Log(input)
			t.Log(e)
			t.Error("did not halt gracefully")
			return
		}

		out = strings.TrimSpace(out)
		output = strings.TrimSpace(output)
		if out != output {
			t.Log(input)
			t.Log("expect: %s", output)
			t.Errorf("got: %s", out)
		}
	}

	o("printInt(3)", "3")
	o("printInt(-3)", "-3")
	o("a := 3; if a == 3 { printInt(5) }", "5")
	o("a := 5; if a == 3 { printInt(5) }", "")
	o("a := 5; if a == 3 { printInt(5) } else { printInt(10) }", "10")
}
