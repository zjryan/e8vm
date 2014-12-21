# thoughts on asm8

the target for a function is basically a bunch of
instructions with possible symbols that can link later

there are many types of symbols
- local labels, these are symbols that are filled based on local function offsets, can be filled when the function is well-defined
- global symbols, these are symbols that are filled on linking time
- when filling a symbol, it might use a set of filling methods

about labels:

labels are indexed positions inside a function