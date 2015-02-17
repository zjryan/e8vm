package build8

import (
	"time"
)

type timeStamp struct {
	t     time.Time
	valid bool
}

func (ts *timeStamp) update(t time.Time) {
	if !ts.valid || t.After(ts.t) {
		ts.t = t
		ts.valid = true
	}
}
