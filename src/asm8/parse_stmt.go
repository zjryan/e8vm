package asm8

import (
	"lex8"
)

type stmt struct {
	*inst
	label string

	ops []*lex8.Token
}

func isValidLabel(s string) bool {
	if len(s) <= 1 || s[0] != '.' {
		return false
	}

	for i, r := range s[1:] {
		if r >= '0' && r <= '9' && i > 0 {
			continue
		}
		if r >= 'a' && r <= 'z' {
			continue
		}
		if r >= 'A' && r <= 'Z' {
			continue
		}
		if r == '_' {
			continue
		}
		return false
	}
	return true
}

func parseStmt(p *Parser) *stmt {
	ops := parseOps(p)
	if len(ops) == 0 {
		return nil
	}

	op0 := ops[0]
	lead := op0.Lit
	if lead == "" {
		panic("empty operand")
	}

	if lead[0] == '.' {
		if !isValidLabel(lead) {
			p.err(op0.Pos, "invalid label")
			return nil
		}
		if len(ops) > 1 {
			p.err(op0.Pos, "label should take the entire line")
			return nil
		}

		return &stmt{label: lead, ops: ops}
	}

	return &stmt{inst: makeInst(p, ops), ops: ops}
}
