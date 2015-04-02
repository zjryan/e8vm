for IR writing, we need to support several operations

- each basic block has exactly one block as its natural next.
- all basic blocks in a function are saved in an order
- if a basic block's natural next is the next in order, then no
  instructions are required at the end
- if a basic block's natural next is not the next in order, then
  an unconditional jump instruction is required at the end

## a basic block split action

- we are writing basic block A, A->B
- we will now add a basic block C
- A->C, C->B
- then we can add conditional jumps at the end of A
- this completes the basic block A


for an if statement
- we are writing basic block A, where A->B
- we finialize A with as A->D, D->B, and if condition A=>C
- then we switch to D to further write the condition expression
- and if met D=>C, etc.
- say the last splitted basic block is X
- we create a basic block C, where C->X
- then we switch to C to write the else body
- and then we switch back the last splitted basic block, and keep
  writing

for while-like for statement
- we are writing basic block A, where A->B
- we create a basic block C, where C->B
