package periodTrigger

import (
	"sync"
	"time"
)

type HashCollection struct {
	locker         sync.RWMutex
	evictDuration  time.Duration
	stableInterval time.Duration

	dict map[string]*periodTrigger
}

func NewHashCollection(evictDuration, stableInterval time.Duration) *HashCollection {
	model := &HashCollection{
		evictDuration:  evictDuration,
		stableInterval: stableInterval,
		dict:           make(map[string]*periodTrigger, 64),
	}
	return model
}

func (c *HashCollection) Check(key string, ts_ns int64, interval time.Duration) TriggerState {
	c.locker.Lock()
	defer c.locker.Unlock()

	v, find := c.dict[key]
	if !find {
		s := newPeriodTrigger(key, interval.Nanoseconds())
		c.dict[key] = s
		resp := s.Check(ts_ns)
		return resp
	} else {
		resp := v.Check(ts_ns)
		return resp
	}
}

func (c *HashCollection) Evict(ts int64) int {
	c.locker.Lock()
	defer c.locker.Unlock()

	inactiveKeys := make([]string, 0, 64)
	for k, v := range c.dict {
		if ts-v.AccessTs >= c.evictDuration.Nanoseconds() {
			inactiveKeys = append(inactiveKeys, k)
		}
	}

	for _, k := range inactiveKeys {
		delete(c.dict, k)
	}
	return len(inactiveKeys)
}
func (c *HashCollection) Count() int {
	c.locker.RLock()
	defer c.locker.RUnlock()

	return len(c.dict)
}
