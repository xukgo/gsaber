package periodTrigger

import "time"

type TriggerState struct {
	Triggered             bool
	IntervalNs            int64
	LastTriggerTs         int64
	LastTriggerTotalCount uint64
	CurrentTotalCount     uint64
}

type IPeriodTriggerCollection interface {
	Check(key string, ts int64, interval time.Duration) TriggerState
	Evict(ts int64) int
	Count() int
}
