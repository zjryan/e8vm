package g8

import (
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
