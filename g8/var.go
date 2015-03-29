package g8

import (
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/sym8"
)

type objVar struct {
	name string
	*ref // the reference of this variable
}

func declareVar(b *builder, tok *lex8.Token, t typ) *ref {
	name := tok.Lit
	v := new(objVar)
	v.name = name
	s := sym8.Make(name, symVar, v, tok.Pos)
	declared := b.scope.Declare(s)
	if declared != nil {
		b.Errorf(tok.Pos, "%q already declared as a %s",
			name, symStr(declared.Type),
		)
		return nil
	}

	// successfuly declared here
	local := b.f.NewLocal(typeSize(t), name)
	v.ref = newRef(t, local)
	return v.ref
}
