package g8

import (
	"bytes"
	"fmt"
)

func typStr(t []typ) string {
	if len(t) == 0 {
		return "<nil>"
	} else if len(t) == 1 {
		return t[0].String()
	}

	buf := new(bytes.Buffer)
	for i, a := range t {
		if i > 0 {
			fmt.Fprint(buf, ",")
		}
		fmt.Fprint(buf, a)
	}
	return buf.String()
}
