package lex8

// Remover removes a particular type of token from a token stream
type Remover struct {
	Tokener
	t int
}

// NewRemover creates a new remover that removes token of type t
func NewRemover(t int) *Remover {
	if t == EOF {
		panic("cannot remove EOF")
	}

	ret := new(Remover)
	ret.t = t

	return ret
}

// Token implements the Tokener interface but only returns
// token that is not the particular type.
func (r *Remover) Token() *Token {
	for {
		ret := r.Token()
		if ret.Type != r.t {
			return ret
		}
	}
}
