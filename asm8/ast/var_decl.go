package ast

import (
	"lonnie.io/e8vm/lex8"
)

type VarDecl struct {
	Stmts []*VarStmt

	Kw, Name             *lex8.Token
	Lbrace, Rbrace, Semi *lex8.Token

	Index uint32 // symbol decl index
}
