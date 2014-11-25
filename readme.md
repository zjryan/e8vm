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

- `arch8/regs.go`: Registers (done)
- `arch8/page.go`: Page (done)
- `arch8/phy_memory.go`: Physical Memory (done)
- `arch8/exception.go`: Exception (done)
- `arch8/page_table.go`: Page Table (done)
- `arch8/virt_memory.go`: Virtual Memory (done)
- `arch8/interrupt.go`: Interrupt Control (done)
- `arch8/cpu.go`: Processor Simulator Structure (done)
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
- `arch8/machine.go`: Bind CPU with some default IO devices. (done)
- `arch8/boot_rom.go`: Boot ROM loader

### Assembler (asm8)

`todo`

### Programming Language (lang8)

`todo`

### Operating System (os8)

`todo`

### Go Language Compiler (go8)

`todo`
