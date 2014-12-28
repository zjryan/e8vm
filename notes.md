# object file

so in asm8, a package should be compiled into an object file
where the syntaxes are checked, but top level symbols are not filled

so a function will look like this:


import packages: name, id

id: symbol id
type: function or var
content: a set of bytes
symbols to fill: a list of (offset, filling method, pack id, symbol id)
	offset and filling method: 32bit
	pack id: 32bit
	sym id: 32bit

should save it in a semi-human readalbe fashion
