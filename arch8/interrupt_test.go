package arch8

import (
	"testing"
)

func TestInterrupt(t *testing.T) {
	as := func(cond bool) {
		if !cond {
			t.Fail()
		}
	}

	p := newPage()

	for i := byte(0); i < 8; i++ {
		in := newInterrupt(p, i)

		has, b := in.Poll()
		as(!has && b == 0)

		in.Enable()
		has, b = in.Poll()
		as(!has && b == 0)

		in.Issue(37)
		has, b = in.Poll()
		as(!has && b == 0)

		in.EnableInt(37)
		has, b = in.Poll()
		as(has && b == 37)

		in.DisableInt(37)
		has, b = in.Poll()
		as(!has && b == 0)

		in.EnableInt(37)
		in.EnableInt(46)
		in.Issue(46)
		has, b = in.Poll()
		as(has && b == 37)

		in.Clear(37)
		has, b = in.Poll()
		as(has && b == 46)

		in.Disable()
		has, b = in.Poll()
		as(!has && b == 0)
	}
}
