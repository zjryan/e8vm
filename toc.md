# Table of Content (Planned)

## Architecture (arch8)

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

## Disassembler (dasm8)

- `dasm8/reg.go`: Register names (done)
- `dasm8/line.go`: Disassemble line structure (done)
- `dasm8/inst_reg.go`: Register instructions (done)
- `dasm8/inst_imm.go`: Immediate instructions (done)
- `dasm8/inst_br.go`: Branch instructions (done)
- `dasm8/inst_sys.go`: System instructions (done)
- `dasm8/inst_jmp.go`: Jump instructions (done)
- `dasm8/inst_all.go`: Put all instructions together (done)
- `dasm8/dasm.go`: Disassembling (done)

## Lexing Framework (lex8)

- `lex8/pos.go`: File position (done)
- `lex8/error.go`: File parsing error (done)
- `lex8/err_list.go`: File parsing error list (done)
- `lex8/rune_scanner.go`: Rune scanner (done)
- `lex8/lex_scanner.go`: Buffered scanner for tokenizing (done)
- `lex8/token.go`: Token structure (done)
- `lex8/lexer.go`: Lexer Framework (done)

## Linker (link8)

- `link8/symbol.go`: Link symbol (done)
- `link8/func.go`: Function, linkable code section (done)
- `link8/var.go`: Var, data section (done)
- `link8/package.go`: Package, building unit (done)
- `link8/pkg_sym.go`: A symbol in a package (done)
- `link8/linker.go`: Linker framework (done)
- `link8/trace_use.go`: Tracing symbol usage (done)
- `link8/layout.go`: Section layout (done)
- `link8/writer.go`: Binary writer (done)
- `link8/link.go`: Link sections together (done)

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
- `ams8/parse_reg.go`: Parsing register (done)
- `asm8/parse_label.go`: Parsing Label (done)
- `asm8/inst.go`: Assembly instruction AST node (done)
- `asm8/stmt.go`: Statement AST node (done)
- `asm8/parse_inst.go`: Assembly instruction parsing framework (done)
- `asm8/inst_reg.go`: Register instructions (done)
- `asm8/inst_imm.go`: Immediate instructions (done)
- `asm8/inst_br.go`: Branch instructions (done)
- `asm8/inst_sys.go`: System instructions (done)
- `asm8/inst_jmp.go`: Jump instructions (done)
- `asm8/inst_all.go`: Put them all together (done)
- `asm8/parse_ops.go`: Parsing operands in a statement (done)
- `asm8/parse_stmt.go`: Assembly instruction and label parsing (done)
- `asm8/parse_func.go`: Function parsing (test done)
- `asm8/parse_var.go`: Variable block parser
- `asm8/parse_const.go`: Const statement parser
- `asm8/parse_import.go`: Import statement parser
- `asm8/parse_file.go`: File parser (partially done)
- `asm8/symbol.go`: Symbol (done)
- `asm8/sym_table.go`: Symbol table (done)
- `asm8/sym_scope.go`: Symbol scope, stack of sym tables (done)
- `asm8/builder.go`: Builder framework (done)
- `asm8/build_func.go`: Function block builder (done)
- `asm8/build_file.go`: File builder (partially done)
- `asm8/build_pkg.go`: Package builder (partially done)
- `asm8/bare_func.go`: Build bare function (done)
- `asm8/single_file.go`: Build single file package function (done)

## Programming Language (lang8)

`todo: a c-like programming language for writing os8`

## Operating System (os8)

`todo`

## Go Language Compiler (go8)

`todo`
