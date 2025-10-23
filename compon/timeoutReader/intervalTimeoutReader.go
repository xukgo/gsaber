package timeoutReader

import (
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
)

type IntervalTimeoutReader struct {
	locker         sync.Mutex
	reader         io.Reader
	interval       time.Duration
	detectInterval time.Duration

	lastReadTime   atomic.Pointer[time.Time]
	timeoutTrigger atomic.Bool
	detectCanceled atomic.Bool
}

func NewIntervalTimeoutReader(r io.Reader, interval time.Duration, detectInterval time.Duration) *IntervalTimeoutReader {
	s := &IntervalTimeoutReader{
		reader:         r,
		interval:       interval,
		detectInterval: detectInterval,
	}
	s.timeoutTrigger.Store(false)
	s.detectCanceled.Store(false)
	t := time.Now()
	s.lastReadTime.Store(&t)
	return s
}

func (c *IntervalTimeoutReader) CheckTimeout() bool {
	return c.timeoutTrigger.Load()
}

func (c *IntervalTimeoutReader) CancelDetect() {
	c.detectCanceled.Store(true)
}

func (c *IntervalTimeoutReader) StartDetect(cb func()) {
	for {
		time.Sleep(c.detectInterval)
		if c.detectCanceled.Load() {
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
	n, err := c.reader.Read(p)
	if err != nil {
		return n, err
	}

	if c.CheckTimeout() {
		return 0, fmt.Errorf("read detect timeout")
	}
	dtNow := time.Now()
	c.lastReadTime.Store(&dtNow)
	return n, err
}
