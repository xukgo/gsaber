package condAction

import (
	"sync"
)

type CountActionWrapper struct {
	locker   sync.Mutex
	interval uint64
	count    uint64
}

func NewCountActionWrapper(interval uint64) *CountActionWrapper {
	w := &CountActionWrapper{
		interval: interval,
		count:    0,
	}
	return w
}

func (c *CountActionWrapper) SetInterval(interval uint64) {
	c.locker.Lock()
	c.interval = interval
	c.locker.Unlock()
}
func (c *CountActionWrapper) SetCount(count uint64) {
	c.locker.Lock()
	c.count = count
	c.locker.Unlock()
}

func (c *CountActionWrapper) Do(action func()) {
	c.locker.Lock()
	c.count++
	if c.count >= c.interval {
		c.count = 0
		c.locker.Unlock()
		action()
	} else {
		c.locker.Unlock()
	}
}
