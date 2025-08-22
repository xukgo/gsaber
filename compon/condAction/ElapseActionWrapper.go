package condAction

import (
	"sync"
	"time"
)

type ElapseActionWrapper struct {
	locker   sync.Mutex
	interval time.Duration
	lastTime time.Time
}

func NewElapseActionWrapper(interval time.Duration, lastTime time.Time) *ElapseActionWrapper {
	w := &ElapseActionWrapper{
		interval: interval,
		lastTime: lastTime,
	}
	return w
}

func (c *ElapseActionWrapper) SetInterval(interval time.Duration) {
	c.locker.Lock()
	c.interval = interval
	c.locker.Unlock()
}
func (c *ElapseActionWrapper) SetLastTime(lastTime time.Time) {
	c.locker.Lock()
	c.lastTime = lastTime
	c.locker.Unlock()
}

func (c *ElapseActionWrapper) Do(currentTime time.Time, action func()) {
	c.locker.Lock()
	if currentTime.Sub(c.lastTime) >= c.interval {
		c.lastTime = currentTime
		c.locker.Unlock()
		action()
	} else {
		c.locker.Unlock()
	}
}
