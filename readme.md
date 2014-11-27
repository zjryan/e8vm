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
- `arch8/inst_reg.go`: Register instructions (done)
- `arch8/inst_imm.go`: Immediate instructions (done)
- `arch8/inst_br.go`: Branch instructions (done)
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

- `asm8/scanner.go`: A line scanner
- `asm8/symbol.go`: Symbol tree
- `asm8/top_parser.go`: Top level syntax parser.
- `asm8/inst.go`: Variable length instruction
- `asm8/func.go`: Function/code section
- `asm8/inst_parser.go`: Instruction parsing helpers
- `asm8/inst_reg.go`: Register instructions
- `asm8/inst_imm.go`: Immediate instructions
- `asm8/inst_br.go`: Branch instructions
- `asm8/inst_jmp.go`: Jump instructions
- `asm8/inst_sys.go`: System instructions
- `asm8/inst_all.go`: Put all instructions together
- `asm8/var.go`: Variable/data section
- `asm8/dat_parse.go`: Data line parser
- `asm8/dat_ints.go`: Hex
- `asm8/dat_str.go`: String
- `asm8/const.go`: Consts
- `asm8/layout.go`: Layout Symbols
- `asm8/build.go`: Build an image

### Programming Language (lang8)

`todo`

### Operating System (os8)

`todo`

### Go Language Compiler (go8)

`todo`
