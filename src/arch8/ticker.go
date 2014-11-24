package arch8

import (
	"math/rand"
	"time"
)

// Ticker is a device that generates time interrupts.
type Ticker struct {
	intBus   IntBus
	nextTick int32

	Interval int32
	Noise    int32
	Rand     *rand.Rand
	Code     byte
}

// NewTicker creates a new time interrupt generator.
func NewTicker(bus IntBus) *Ticker {
	ret := new(Ticker)
	ret.intBus = bus

	ret.Interval = 1000
	ret.Noise = 10
	ret.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	ret.Code = 2 // time interrupt code

	ret.reset()

	return ret
}

func (t *Ticker) reset() {
	if t.Noise < 0 {
		panic("negative ticker noise")
	}
	if t.Interval < 0 {
		panic("negative ticker interval")
	}

	noise := int32(0)
	if t.Noise > 0 {
		noise = t.Rand.Int31n(t.Noise) - t.Noise/2
	}

	next := t.Interval + noise
	if next < 0 {
		next = 0
	}

	t.nextTick = next
}

// Tick decreases the ticking counter. If the counter reaches 0,
// it will issue interrupts to all cores and reset the counter.
func (t *Ticker) Tick() {
	if t.nextTick < 0 {
		panic("bug")
	}

	if t.nextTick == 0 {
		IntAllCores(t.intBus, t.Code)
		t.reset()
	} else {
		t.nextTick--
	}
}
