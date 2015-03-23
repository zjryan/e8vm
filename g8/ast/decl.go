package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Decl: var, function, struct or interface
type Decl interface{}

type IdentList struct {
	Idents []*lex8.Token
	Commas []*lex8.Token
}

// VarDecl declarse a set of variable
type VarDecl struct {
	Idents *IdentList
	Type   Expr
	Eq     *lex8.Token
	Exprs  *ExprList
	Semi   *lex8.Token
}

// ConstDecl declares a set of constants
type ConstDecl struct {
	Ident *IdentList
	Type  Expr
	Eq    *lex8.Token
	Exprs *ExprList
	Semi  *lex8.Token
}

// ConstDecls is a const declaration with a leading keyword
// It could be a single decl or a decl block.
type ConstDecls struct {
	Kw     *lex8.Token
	Lparen *lex8.Token // optional
	Decls  []*ConstDecl
	Rparen *lex8.Token // optional
	Semi   *lex8.Token
}

// VarDecls is a variable declaration with a leading keyword
// It could be a single decl or a decl block.
type VarDecls struct {
	Kw     *lex8.Token
	Lparen *lex8.Token // optional
	Decls  []*VarDecl
	Rparen *lex8.Token // optional
	Semi   *lex8.Token
}
