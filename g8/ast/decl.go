package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Decl is a general declaration,
// could be var, function, struct or interface
type Decl interface{}

// IdentList is a list of identifiers
type IdentList struct {
	Idents []*lex8.Token
	Commas []*lex8.Token
}

// Para is a function parameter
type Para struct {
	Ident *lex8.Token
	Type  Expr // when Type is missing, Ident also might be the type
}

// ParaList is a parameter list
type ParaList struct {
	Paras  []*Para
	Commas []*lex8.Token
}

// FuncDecl declares a function
type FuncDecl struct {
	Kw   *lex8.Token
	Name *lex8.Token

	Lparen *lex8.Token
	Args   *ParaList
	Rparen *lex8.Token

	// ret list
	RetLparen *lex8.Token // optional
	Rets      *ParaList
	RetRparen *lex8.Token // optional

	// single ret type
	RetType Expr

	Body *Block
	Semi *lex8.Token
}

// VarDecl declares a set of variable
type VarDecl struct {
	Idents *IdentList
	Type   Expr
	Eq     *lex8.Token
	Exprs  *ExprList
	Semi   *lex8.Token
}

// Field is a member variable of a struct
type Field struct {
	Idents *IdentList
	Type   Expr
}

// StructDecl declarse a structure type
type StructDecl struct {
	Kw     *lex8.Token
	Name   *lex8.Token
	Lbrace *lex8.Token

	Fields  []*Field
	Methods []*FuncDecl

	Rbrace *lex8.Token
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
