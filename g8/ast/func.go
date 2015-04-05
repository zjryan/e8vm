package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Para is a function parameter
type Para struct {
	Ident *lex8.Token
	Type  Expr // when Type is missing, Ident also might be the type
}

// ParaList is a parameter list
type ParaList struct {
	Lparen *lex8.Token
	Paras  []*Para
	Commas []*lex8.Token
	Rparen *lex8.Token
}

// FuncSig is a function Signature
type FuncSig struct {
	Name    *lex8.Token
	Args    *ParaList
	Rets    *ParaList // ret list
	RetType Expr      // single ret type
}

// Func is a function
type Func struct {
	Kw   *lex8.Token
	Name *lex8.Token

	*FuncSig

	Body *Block
	Semi *lex8.Token
}
