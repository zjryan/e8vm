package arch8

// Serial is a serial console device
// It is basically two DMA pipes of ring buffered bytes:
// one pipe for input, one pipe for output.
type Serial struct {
	p *Page
}

const (
	serialBase     = 128
	serialInHead   = serialBase + 0  // updated by hardware
	serialInTail   = serialBase + 4  // updated by cpu
	serialInWait   = serialBase + 8  // cycles to wait to raise an interrupt
	serialInThresh = serialBase + 12 // threashold to raise an input interrupt

	serialOutHead   = serialBase + 16 // updated by cpu
	serialOutTail   = serialBase + 20 // updated by hardware
	serialOutWait   = serialBase + 24 // cycles to wait to raise an interrupt
	serialOutThresh = serialBase + 28 // threashold to raise an output interrupt

	serialInputBuf  = serialBase + 64
	serialOutputBuf = serialBase + 96

	serialCap = 30 // 30 bytes maximum in each pipe
)
