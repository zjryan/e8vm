package parse

import (
	"fmt"
	"io"

	"lonnie.io/e8vm/lex8"
)

type parser struct {
	x lex8.Tokener
	*lex8.Parser
}

func makeTokener(f string, r io.Reader) (lex8.Tokener, *lex8.Recorder) {
	var x lex8.Tokener = newLexer(f, r)
	x = newSemiInserter(x)
	rec := lex8.NewRecorder(x)
	return lex8.NewCommentRemover(rec), rec
}

func newParser(f string, r io.Reader) (*parser, *lex8.Recorder) {
	ret := new(parser)
	x, rec := makeTokener(f, r)
	ret.x = x
	ret.Parser = lex8.NewParser(ret.x, Types)
	return ret, rec
}

func (p *parser) SeeOp(ops ...string) bool {
	t := p.Token()
	if t.Type != Operator {
		return false
	}
	for _, op := range ops {
		if t.Lit == op {
			return true
		}
	}
	return false
}

func (p *parser) typeStr(t *lex8.Token) string {
	if t.Type == Operator {
		return fmt.Sprintf("'%s'", t.Lit)
	} else if t.Type == Semi {
		return "';'"
	}
	return TypeStr(t.Type)
}

func (p *parser) AcceptSemi() *lex8.Token {
	if p.InError() {
		return nil
	}

	t := p.Token()
	if t.Type == Operator && (t.Lit == "}" || t.Lit == ")") {
		return t // fake semicolon by operator
	}

	if t.Type != Semi {
		return nil
	}
	return p.Shift()
}

func (p *parser) SeeSemi() bool {
	t := p.Token()
	if t.Type == Semi {
		return true
	}
	if t.Type == Operator && (t.Lit == "}" || t.Lit == ")") {
		return true
	}
	return false
}

func (p *parser) ExpectSemi() *lex8.Token {
	if p.InError() {
		return nil
	}

	t := p.Token()
	if t.Type == Operator && (t.Lit == "}" || t.Lit == ")") {
		return t // fake semicolon by operator
	}

	if t.Type != Semi {
		p.ErrorfHere("expect ';', got %s", p.typeStr(t))
		return nil
	}
	return p.Shift()
}

func (p *parser) skipErrStmt() bool {
	if !p.InError() {
		return false
	}

	for {
		t := p.Token()
		if t.Type == Semi || t.Type == lex8.EOF {
			break
		} else if t.Type == Operator && t.Lit == "}" {
			break
		}
		p.Next()
	}
	if p.See(Semi) {
		p.Next()
	}

	p.BailOut()
	return true
}

func (p *parser) SeeKeyword(kw string) bool {
	return p.SeeLit(Keyword, kw)
}

func (p *parser) ExpectOp(op string) *lex8.Token {
	if p.InError() {
		return nil
	}
	t := p.Token()
	if t.Type != Operator || t.Lit != op {
		p.ErrorfHere("expect '%s', got %s", op, p.typeStr(t))
		return nil
	}

	return p.Shift()
}

func (p *parser) ExpectKeyword(kw string) *lex8.Token {
	if !p.SeeLit(Keyword, kw) {
		p.ErrorfHere("expect keyword '%s', got %s",
			kw, p.typeStr(p.Token()),
		)
		return nil
	}
	return p.Shift()
}

// Tokens parses a file into a token array
func Tokens(f string, r io.Reader) ([]*lex8.Token, []*lex8.Error) {
	x, _ := makeTokener(f, r)
	toks := lex8.TokenAll(x)
	if errs := x.Errs(); errs != nil {
		return nil, errs
	}
	return toks, nil
}
