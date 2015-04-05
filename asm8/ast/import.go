package ast

import (
	"lonnie.io/e8vm/lex8"
)

// Import is an import declaration block
type Import struct {
	Stmts []*ImportStmt

	Kw, Lbrace, Rbrace, Semi *lex8.Token
}

// ImportStmt is an import statement
type ImportStmt struct {
	Path *lex8.Token
	As   *lex8.Token
}
