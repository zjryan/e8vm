# E8VM

Goal: A book written in working code and markdown document that
describes how computer system works, including architecture,
assemblers, compilers, and operating systems.

Planned Features:

- **Modularized.** File based modules. No circular dependency (not only on packages,
  but also on files). A reader can always read the project a file by
  a file, either from bottom to top, or from top to bottom.
- **Small files.** Each file is shorter than 200 lines of code.
- **Tested and Documented.**
  Each file (will) come with test cases, examples, and markdown description.
- **Real.** The simulation (will) work like a real computer.

## Table of Content (Planned)

###  Architecture (arch8)

- `arch8/regs.go`: Registers (test done)
- `arch8/page.go`: Page (test done)
- `arch8/phy_memory.go`: Physical Memory (test done)
- `arch8/exception.go`: Exception (done, no test)
- `arch8/page_table.go`: Page Table (test done)
- `arch8/virt_memory.go`: Virtual Memory (test done)
- `arch8/interrupt.go`: Interrupt Control (test done)
- `arch8/cpu.go`: Processor Simulator Structure (test done)
- `arch8/inst_reg.go`: Register instructions (test done)
- `arch8/inst_imm.go`: Immediate instructions (test done)
- `arch8/inst_br.go`: Branch instructions (test done)
- `arch8/inst_jmp.go`: Jump instructions (done)
- `arch8/inst_sys.go`: System instructions (done)
- `arch8/inst_all.go`: Put all instructions together (done)
- `arch8/int_bus.go`: Interrupt bus (done)
- `arch8/multi_core.go`: Shared memory multicore (done)
- `arch8/device.go`: General IO device (done)
- `arch8/serial.go`: Serial Console Control (done)
- `arch8/ticker.go`: A ticker that generates time interrupts. (done)
- `arch8/machine.go`: Bind stuff together and image loading. (done)
- `arch8/run_raw.go`: A shortcut function to run a raw image. (done)

### Assembler (asm8)

- `asm8/pos.go`: File position (done)
- `asm8/error.go`: File parsing error (done)
- `asm8/err_list.go`: File parsing error list (done)
- `asm8/rune_scanner.go`: Rune scanner (done)
- `asm8/lex_scanner.go`: Buffered scanner for tokenizing (done)
- `asm8/token.go`: Asm8 Tokens (done)
- `asm8/lexer.go`: Lexer Framework (done)
- `asm8/lex_comment.go`: Lexing comments (done)
- `asm8/lex_string.go`: Lexing strings (done)
- `asm8/lex_operand.go`: Lexing operands (done)
- `asm8/lex_all.go`: Put the lexing altogether (test done)
- `asm8/parser.go`: Parser framework (done)
- `asm8/parse_func.go`: Function block parser (test done)
- `asm8/parse_var.go`: Variable block parser
- `asm8/parse_const.go`: Const statement parser
- `asm8/parse_import.go`: Import statement parser
- `asm8/parse_all.go`: Put the parsing altogether
- `asm8/builder.go`: Builder framework
- `asm8/build_func.go`: Function block builder
- `asm8/sym_scope.go`: Symbol scope
- `asm8/sym_table.go`: Symbol table, stack of symbol scopes
- `asm8/inst.go`: Variable length instruction
- `asm8/inst_parser.go`: Instruction parsing helpers
- `asm8/inst_reg.go`: Register instructions
- `asm8/inst_imm.go`: Immediate instructions
- `asm8/inst_br.go`: Branch instructions
- `asm8/inst_jmp.go`: Jump instructions
- `asm8/inst_sys.go`: System instructions
- `asm8/inst_all.go`: Put all instructions together
- `asm8/layout.go`: Layout symbols
- `asm8/image.go`: Image builder

### Programming Language (lang8)

`todo: a c-like programming language for writing os8`

### Operating System (os8)

`todo`

### Go Language Compiler (go8)

`todo`
