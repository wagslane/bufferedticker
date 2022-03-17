package bufferedticker

import "time"

// Ticker -
type Ticker struct {
	originalTicker *time.Ticker
	C              <-chan time.Time
}

// NewTicker returns a new buffered Ticker containing a channel that will send the current time on the channel after each tick.
// The period of the ticks is specified by the duration argument.
// The ticker will adjust the time interval or drop ticks to make up for slow receivers.
// The duration d must be greater than zero; if not, NewTicker will panic. Stop the ticker to release associated resources.
func NewTicker(d time.Duration, bufferedTicksCount int) Ticker {
	bufferedTicker := make(chan time.Time, bufferedTicksCount)
	originalTicker := time.NewTicker(d)

	go func() {
		for tick := range originalTicker.C {
			bufferedTicker <- tick
		}
	}()

	return Ticker{
		originalTicker: originalTicker,
		C:              bufferedTicker,
	}
}

// Reset resets its period to the specified duration.
// The next new tick will arrive after the new period elapses, but all buffered ticks
// will be unaffected.
// The duration d must be greater than zero; if not, Reset will panic.
func (t *Ticker) Reset(d time.Duration) {
	t.originalTicker.Reset(d)
}

// Stop turns off a ticker.
// After Stop, no more ticks will be buffered.
// Stop does not close the channel, to prevent a concurrent goroutine reading from the channel from seeing an erroneous "tick".
func (t *Ticker) Stop() {
	t.originalTicker.Stop()
}
