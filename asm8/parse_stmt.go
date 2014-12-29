package asm8

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

	if parseLabel(p, op0) {
		if len(ops) > 1 {
			p.err(op0.Pos, "label should take the entire line")
			return nil
		}
		return &stmt{label: lead, ops: ops}
	}

	return &stmt{inst: parseInst(p, ops), ops: ops}
}

func (s *stmt) isLabel() bool {
	return s.inst == nil && s.label != ""
}
