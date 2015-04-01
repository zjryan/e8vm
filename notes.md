build in functions are assembly functions with a high-level language signature.
0 is the package itself.

- we should a type of expression list. It is not a type of type that
  you can declare, but it is a internal type for assignment and stuff

here is the issue:
- a list of expression is a type on assigning and as the return value
- of a call expression.

for:

func f(x, y int) (b, c int) 
it is valid to do this:

- x, y = f(a,b)
- x, y := f(a,b)
- f(f(a,b))

but it is invalid to do this:
- f(f(a,b), c) // cannot recursively composite an expression list
- f(a,b) + c // cannot perform arith ops on expression list

as a result, even we add expression list as a type,
we need to add checkings at many places.
