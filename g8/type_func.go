package g8

import (
	"bytes"
	"fmt"

	"lonnie.io/e8vm/fmt8"
	"lonnie.io/e8vm/g8/ir"
)

type typFunc struct {
	argTypes []typ
	retTypes []typ

	// optional names
	argNames []string
	retNames []string

	sig *ir.FuncSig // caching the IR sig
}

func (t *typFunc) String() string {
	// TODO: this is kind of ugly, need some refactor
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "func (%s) ", fmt8.Join(t.argTypes, ","))
	if len(t.retTypes) > 1 {
		fmt.Fprintf(buf, "(")
		for i, ret := range t.retTypes {
			if i > 0 {
				fmt.Fprint(buf, ",")
			}
			fmt.Fprint(buf, ret)
		}
		fmt.Fprint(buf, ")")
	} else if len(t.retTypes) == 1 {
		fmt.Fprint(buf, t.retTypes[0])
	}

	return buf.String()
}

func (t *typFunc) Size() int32 { return 4 }
