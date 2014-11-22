(plan and progress)

- arch8: The Architecture
    - regs.go: Registers (done)
    - page.go: Page (done)
    - phy_memory.go: Physical Memory (done)
    - page_table.go: Page Table
    - virt_memory.go: Vitural Memory // do we really need page table? 
    - interrupt.go 
    - inst_reg.go: Register instructions
    - inst_im.go: Immediate instructions
    - inst_b.go: Branch instructions
    - inst_j.go: Jump instructions
    - inst_x.go: Other instructions
    - syscall.go: System Calls
    - core.go: Processor Simulator
    - multi.go: Multicore Simulator
    - io.go: IO devices
    - rom.go: Bootloader

- asm8: The Assembler
    - symbol.go: Symbol system
    - inst.go: Instruction with symbols
    - inst_stmt.go: Single Instruction Parsing
    - func.go: Code segment encoding
    - var.go: Data segment encoding
    - const.go: Constant number
    - parse.go: Asm file parsing
    - link.go: Asm file linking

- g8: The Language
    - g8.go: A C like language but with namespaces

- os8: The Operating system
    - os8.go: An operating system that supports threads and processes
