package arch8

// Interrupt defines the interrupt page
type Interrupt struct {
	p *Page // the dma page for interrupt handler
}

// 256 interrupts, has priorities, but no preemptive 
// enable switch (1 bit) - master interrupt enabling switch
// current interrupt (1 byte)
// enable masks (32 bytes) - set when interrupt is enabled 
// pending masks (32 bytes) - set by hardward when the interrupt is pending
// saved contexts (32 bytes) - the registers and the interrupt number
// interrupt handlers (1024 bytes) - interrupt handler pc