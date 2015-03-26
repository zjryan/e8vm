package ir

// Main is just an entrance for simple testing
func Main() {
	p := newPkg("_")
	f := p.newFunc()
	b := f.newBlock()

	v1 := f.newTemp(regSize)
	v2 := f.newTemp(regSize)

	b.assign(v1, snum(3))
	b.assign(v2, snum(4))
	b.arith(v1, v1, "+", v2)

	genPkg(p)
}
