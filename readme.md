# E8VM

Goal: A book written in working code and markdown document that
describes how computer system works, including architecture,
assemblers, compilers, and operating systems.

## Table of Content (Planned)

###  Architecture (arch8)

- `arch8/regs.go`: Registers (done)
- `arch8/page.go`: Page (done)
- `arch8/phy_memory.go`: Physical Memory (done)
- `arch8/exception.go`: Exception (done)
- `arch8/page_table.go`: Page Table (done)
- `arch8/virt_memory.go`: Virtual Memory (done)
- `arch8/interrupt.go`: Interrupt Control (done)
- `arch8/serial.go`: Serial Console Control (done)
- `arch8/cpu.go`: Processor Simulator Structure (done)
- `arch8/inst_reg.go`: Register instructions
- `arch8/inst_imm.go`: Immediate instructions
- `arch8/inst_br.go`: Branch instructions
- `arch8/inst_jmp.go`: Jump instructions
- `arch8/inst_sys.go`: System instructions
- `arch8/inst.go`: Put all the instructions altogether
- `arch8/multi.go`: Shared memory multicore
- `arch8/boot_rom.go`: Boot ROM loader

### Assembler (asm8)

- `asm8/symbol.go`: Symbol system

### Programming Language (lang8)

- `lang8/lang8.go`: Introduction

### Operating System (os8)

- `os8/os8.go`: Introduction
