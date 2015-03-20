package ast

import (
	"lonnie.io/e8vm/lex8"
)

// ImportDecl is an import declaration block
type ImportDecl struct {
	Stmts []*ImportStmt

	Kw, Lbrace, Rbrace, Semi *lex8.Token
}

// ImportStmt is an import statement
type ImportStmt struct {
	Path *lex8.Token
	As *lex8.Token
}
