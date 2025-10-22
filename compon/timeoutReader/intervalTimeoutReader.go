package timeoutReader

import (
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

type IntervalTimeoutReader struct {
	reader         io.Reader
	interval       time.Duration
	detectInterval time.Duration

	//timer          *time.Timer
	lastReadTime   atomic.Pointer[time.Time]
	timeoutTrigger atomic.Bool
	detectCanceld  atomic.Bool
}

func NewIntervalTimeoutReader(r io.Reader, interval time.Duration, detectInterval time.Duration) *IntervalTimeoutReader {
	s := &IntervalTimeoutReader{
		reader:         r,
		interval:       interval,
		detectInterval: detectInterval,
		//timer:    time.NewTimer(interval),
	}
	s.timeoutTrigger.Store(false)
	s.detectCanceld.Store(false)
	t := time.Now()
	s.lastReadTime.Store(&t)
	return s
}

func (c *IntervalTimeoutReader) CheckTimeout() bool {
	return c.timeoutTrigger.Load()
}

func (c *IntervalTimeoutReader) CancelDetect() {
	c.detectCanceld.Store(true)
}

func (c *IntervalTimeoutReader) StartIntervalDetect(cb func()) {
	for {
		time.Sleep(c.detectInterval)
		if c.detectCanceld.Load() {
			return
		}

		dtNow := time.Now()
		t := c.lastReadTime.Load()

		if dtNow.Sub(*t).Milliseconds() >= c.interval.Milliseconds() {
			c.timeoutTrigger.Store(true)
			if cb != nil {
				cb()
			}
			return
		}
	}
}

func (c *IntervalTimeoutReader) Read(p []byte) (int, error) {
	if c.CheckTimeout() {
		return 0, fmt.Errorf("read detect timeout")
	}

	n, err := c.reader.Read(p)
	if c.CheckTimeout() {
		return 0, fmt.Errorf("read detect timeout")
	}

	dtNow := time.Now()
	c.lastReadTime.Store(&dtNow)
	return n, err
}
