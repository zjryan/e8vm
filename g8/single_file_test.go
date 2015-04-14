package g8

import (
	"testing"

	"strings"

	"lonnie.io/e8vm/arch8"
)

func TestSingleFile_good(t *testing.T) {
	const N = 100000

	o := func(input, output string) {
		bs, es, _ := CompileSingleFile("main.g", input)
		if es != nil {
			t.Log(input)
			for _, e := range es {
				t.Log(e)
			}
			t.Error("compile failed")
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

	o(`
		func fabo(n int) int {
			if n == 0 return 0
			if n == 1 return 1
			return fabo(n-1) + fabo(n-2)
		}
		func main() { printInt(fabo(10)) }
	`, "55")

	o("func main() { printInt(3) }", "3")
}
