some expressions are also type

- identifier: A, might be a type, might be an expression
- star: *A, might be a pointer type, might be an expression

also, some expressions starts with a type
- []int is a type
- []int { 3, 4, 5 } is an expression
- A might be a type
- A { 3, 5, 7 } is an expression, A must be a type
- A(something) is an expression, this is conversion or 

type is just a special kind of expression
we can have a parse type function, that returns an expression, which
only accepts type.
but for the parse expression function, we have parse the types again



