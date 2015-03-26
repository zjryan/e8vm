package g8

import (
	"lonnie.io/e8vm/lex8"
)

type builder struct {
	*lex8.ErrorList
}

func newBuilder() *builder {
	ret := new(builder)
	ret.ErrorList = lex8.NewErrorList()
	return ret
}
