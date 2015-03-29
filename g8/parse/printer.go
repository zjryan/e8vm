package parse

import (
	"bytes"
	"fmt"
	"strings"
)

type printer struct {
	buf     *bytes.Buffer
	midLine bool

	indent    int
	indentStr string
}

func newPrinter() *printer {
	ret := new(printer)
	ret.buf = new(bytes.Buffer)
	ret.indentStr = "    "
	return ret
}

func (p *printer) printStr(s string) {
	if strings.HasSuffix(s, "\n") {
		p.midLine = false
	} else if s != "" {
		if !p.midLine {
			// print indent before printing the line
			for i := 0; i < p.indent; i++ {
				fmt.Fprint(p.buf, p.indentStr)
			}
		}
		p.midLine = true
	}
	fmt.Fprint(p.buf, s)
}

func (p *printer) printEndl() { p.printStr("\n") }
func (p *printer) Tab()       { p.indent++ }
func (p *printer) ShiftTab() {
	if p.indent > 0 {
		p.indent--
	}
}

func (p *printer) String() string { return p.buf.String() }
