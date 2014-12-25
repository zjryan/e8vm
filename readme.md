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

### Architecture (arch8)

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
- `arch8/inst_sys.go`: System instructions (done)
- `arch8/inst_jmp.go`: Jump instructions (done)
- `arch8/inst_all.go`: Put all instructions together (done)
- `arch8/int_bus.go`: Interrupt bus (done)
- `arch8/multi_core.go`: Shared memory multicore (done)
- `arch8/device.go`: General IO device (done)
- `arch8/serial.go`: Serial Console Control (done)
- `arch8/ticker.go`: A ticker that generates time interrupts. (done)
- `arch8/machine.go`: Bind stuff together and image loading. (done)
- `arch8/run_raw.go`: A shortcut function to run a raw image. (done)

### Disassembler (dasm8)

- `dasm8/reg.go`: Register names (done)
- `dasm8/line.go`: Disassemble line structure (done)
- `dasm8/inst_reg.go`: Register instructions (done)
- `dasm8/inst_imm.go`: Immediate instructions (done)
- `dasm8/inst_br.go`: Branch instructions (done)
- `dasm8/inst_sys.go`: System instructions (done)
- `dasm8/inst_jmp.go`: Jump instructions (done)
- `dasm8/inst_all.go`: Put all instructions together (done)
- `dasm8/dasm.go`: Disassembling (done)

### Lexing Framework (lex8)

- `lex8/pos.go`: File position (done)
- `lex8/error.go`: File parsing error (done)
- `lex8/err_list.go`: File parsing error list (done)
- `lex8/rune_scanner.go`: Rune scanner (done)
- `lex8/lex_scanner.go`: Buffered scanner for tokenizing (done)
- `lex8/token.go`: Token structure (done)
- `lex8/lexer.go`: Lexer Framework (done)

### Assembler (asm8)

- `asm8/token.go`: Asm8 Tokens (done)
- `asm8/lex_comment.go`: Lexing comments (done)
- `asm8/lex_string.go`: Lexing strings (done)
- `asm8/lex_operand.go`: Lexing operands (done)
- `asm8/lex_all.go`: Put the lexing altogether (test done)
- `asm8/stmt_lexer.go`: Replace endl with semicolons (done)
- `asm8/parser.go`: Parser framework (done)
- `asm8/parse_reg.go`: Register names (done)
- `asm8/parse_label.go`: Label name check (done)
- `asm8/parse_sym.go`: Symbol name check (done)
- `asm8/parse_arg.go`: Argument count check (done)
- `asm8/inst.go`: Assembly instruction line datas tructure (done)
- `asm8/parse_inst.go`: Assembly instruction parsing framework (done)
- `asm8/inst_reg.go`: Register instructions (done)
- `asm8/inst_imm.go`: Immediate instructions (done)
- `asm8/inst_br.go`: Branch instructions (done)
- `asm8/inst_sys.go`: System instructions (done)
- `asm8/inst_jmp.go`: Jump instructions (done)
- `asm8/inst_all.go`: Put them all together (done)
- `asm8/parse_stmt.go`: Assembly instruction and label parsing (done)
- `asm8/parse_func.go`: Function parsing (test done)
- `asm8/parse_var.go`: Variable block parser
- `asm8/parse_const.go`: Const statement parser
- `asm8/parse_import.go`: Import statement parser
- `asm8/parse_all.go`: Put the parsing altogether
- `asm8/sym_table.go`: Symbol table (done)
- `asm8/sym_scope.go`: Symbol scope, stack of sym tables (done)
- `asm8/resolve.go`: Symbol resolve check
- `asm8/builder.go`: Builder framework
- `asm8/build_func.go`: Function block builder
- `asm8/link.go`: Layout and linking 
- `asm8/image.go`: Image builder

### Programming Language (lang8)

`todo: a c-like programming language for writing os8`

### Operating System (os8)

`todo`

### Go Language Compiler (go8)

`todo`
