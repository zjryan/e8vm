package parse

import (
	"lonnie.io/e8vm/lex8"
)

// StmtLexer replaces end-lines with semicolons
type semiInserter struct {
	x          lex8.Tokener
	save       *lex8.Token
	insertSemi bool
}

// newSemiInserter creates a new statement lexer that inserts
// semicolons into a token stream.
func newSemiInserter(x lex8.Tokener) *semiInserter {
	ret := new(semiInserter)
	ret.x = x

	return ret
}

// Token returns the next token of lexing
func (sx *semiInserter) Token() *lex8.Token {
	if sx.save != nil {
		ret := sx.save
		sx.save = nil
		return ret
	}

	for {
		t := sx.x.Token()
		switch t.Type {
		case Lbrace, Semi:
			sx.insertSemi = false
		case lex8.EOF:
			if sx.insertSemi {
				sx.insertSemi = false
				sx.save = t
				return &lex8.Token{
					Type: Semi,
					Lit:  t.Lit,
					Pos:  t.Pos,
				}
			}
		case Rbrace:
			if sx.insertSemi {
				sx.save = t
				return &lex8.Token{
					Type: Semi,
					Lit:  t.Lit,
					Pos:  t.Pos,
				}
			}
			sx.insertSemi = true
		case Endl:
			if sx.insertSemi {
				sx.insertSemi = false
				return &lex8.Token{
					Type: Semi,
					Lit:  "\n",
					Pos:  t.Pos,
				}
			}
			continue // ignore this end line
		case lex8.Comment:
			// do nothing
		default:
			sx.insertSemi = true
		}

		return t
	}
}

// Errs returns the list of lexing errors.
func (sx *semiInserter) Errs() []*lex8.Error {
	return sx.x.Errs()
}
