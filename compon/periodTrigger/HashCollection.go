package periodTrigger

import (
	"sync"
	"time"

	"github.com/xukgo/gsaber/utils/reflectUtil"
)

type HashCollection struct {
	evictDuration  time.Duration
	stableInterval time.Duration

	strLocker  sync.RWMutex
	strDict    map[string]*strPeriodTrigger
	uintLocker sync.RWMutex
	uintDict   map[uint64]*uintPeriodTrigger
}

func NewHashCollection(evictDuration, stableInterval time.Duration) *HashCollection {
	model := &HashCollection{
		evictDuration:  evictDuration,
		stableInterval: stableInterval,
		strDict:        make(map[string]*strPeriodTrigger, 64),
		uintDict:       make(map[uint64]*uintPeriodTrigger, 64),
	}
	return model
}

func (c *HashCollection) CheckDefaultFuncLine(ts_ns int64, interval time.Duration) TriggerState {
	key := reflectUtil.FormatCallerLineKey(1)
	return c.Check(key, ts_ns, interval)
}

func (c *HashCollection) Check(key string, ts_ns int64, interval time.Duration) TriggerState {
	c.strLocker.Lock()
	defer c.strLocker.Unlock()

	v, find := c.strDict[key]
	if !find {
		s := newStrPeriodTrigger(key, interval.Nanoseconds())
		c.strDict[key] = s
		resp := s.Check(ts_ns)
		return resp
	} else {
		resp := v.Check(ts_ns)
		return resp
	}
}

func (c *HashCollection) CheckUint(key uint64, ts_ns int64, interval time.Duration) TriggerState {
	c.strLocker.Lock()
	defer c.strLocker.Unlock()

	v, find := c.uintDict[key]
	if !find {
		s := newUintPeriodTrigger(key, interval.Nanoseconds())
		c.uintDict[key] = s
		resp := s.Check(ts_ns)
		return resp
	} else {
		resp := v.Check(ts_ns)
		return resp
	}
}

func (c *HashCollection) Evict(ts int64) int {
	c.strLocker.Lock()
	inactiveKeys1 := make([]string, 0, 4)
	for k, v := range c.strDict {
		if ts-v.AccessTs >= c.evictDuration.Nanoseconds() {
			inactiveKeys1 = append(inactiveKeys1, k)
		}
	}
	for _, k := range inactiveKeys1 {
		delete(c.strDict, k)
	}
	c.strLocker.Unlock()

	c.uintLocker.Lock()
	inactiveKeys2 := make([]uint64, 0, 4)
	for k, v := range c.uintDict {
		if ts-v.AccessTs >= c.evictDuration.Nanoseconds() {
			inactiveKeys2 = append(inactiveKeys2, k)
		}
	}
	for _, k := range inactiveKeys2 {
		delete(c.uintDict, k)
	}
	c.uintLocker.Unlock()

	return len(inactiveKeys1) + len(inactiveKeys2)
}

func (c *HashCollection) Count() int {
	c.strLocker.RLock()
	c.uintLocker.RLock()
	count := len(c.strDict) + len(c.uintDict)
	c.strLocker.RUnlock()
	c.uintLocker.RUnlock()
	return count
}
