package asm8

import (
	"strings"

	"lonnie.io/e8vm/asm8/parse"
	"lonnie.io/e8vm/lex8"
	"lonnie.io/e8vm/link8"
)

// Symbol is a data structure for saving a symbol.
type symbol struct {
	Name    string
	Type    int
	Item    interface{}
	Pos     *lex8.Pos
	Package string // Package path
}

func (s *symbol) clone() *symbol {
	return &symbol{s.Name, s.Type, s.Item, s.Pos, s.Package}
}

// Symbol types
const (
	SymNone   = iota
	SymFunc   // Item.type == *Func
	SymVar    // Item.type == *Var
	SymConst  // Item.type == *Const // TODO
	SymImport // Item.type == *PkgImport
	SymLabel  // Item.type == *stmt
)

func init() {
	as := func(b bool) {
		if !b {
			panic("bug")
		}
	}
	as(SymNone == link8.SymNone)
	as(SymConst == link8.SymConst)
	as(SymFunc == link8.SymFunc)
	as(SymVar == link8.SymVar)
}

func symStr(s int) string {
	switch s {
	case SymImport:
		return "import"
	case SymFunc:
		return "function"
	case SymConst:
		return "constant"
	case SymVar:
		return "variable"
	case SymLabel:
		return "label"
	}
	return "unknown"
}

// IsPublic checks if a symbol name is public.
func isPublic(name string) bool {
	if name == "" {
		return false
	}
	r := name[0]
	return r >= 'A' && r <= 'Z'
}

// mightBeSymbol just looks at the first rune and see if it is *poosibly* a symbol
func mightBeSymbol(sym string) bool {
	if sym == "" {
		return false
	}
	r := sym[0]
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func parseSym(p lex8.Logger, t *lex8.Token) (pack, sym string) {
	if t.Type != parse.Operand {
		panic("symbol not an operand")
	}

	sym = t.Lit
	dot := strings.Index(sym, ".")
	if dot >= 0 {
		pack, sym = sym[:dot], sym[dot+1:]
	}

	if dot >= 0 && !lex8.IsPkgName(pack) {
		p.Errorf(t.Pos, "invalid package name: %q", pack)
	} else if !parse.IsIdent(sym) {
		p.Errorf(t.Pos, "invalid symbol: %q", t.Lit)
	}

	return pack, sym
}
