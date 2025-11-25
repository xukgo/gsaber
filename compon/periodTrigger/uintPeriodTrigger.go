package periodTrigger

type uintPeriodTrigger struct {
	Key                   uint64
	IntervalNs            int64 // unit:ns
	AccessTs              int64 // unit:ns
	LastTriggerTs         int64 // unit:ns
	LastTriggerTotalCount uint64
	CurrentTotalCount     uint64
}

func newUintPeriodTrigger(key uint64, intervalNs int64) *uintPeriodTrigger {
	model := new(uintPeriodTrigger)
	model.Key = key
	model.IntervalNs = intervalNs
	model.AccessTs = 0
	model.LastTriggerTs = 0
	model.LastTriggerTotalCount = 0
	model.CurrentTotalCount = 0
	return model
}

func (c *uintPeriodTrigger) Check(ts int64) TriggerState {
	result := TriggerState{}
	result.IntervalNs = c.IntervalNs
	result.LastTriggerTs = c.LastTriggerTs
	result.LastTriggerTotalCount = c.LastTriggerTotalCount

	c.AccessTs = ts
	c.CurrentTotalCount++
	result.CurrentTotalCount = c.CurrentTotalCount

	if c.CurrentTotalCount > 1 && ts-c.LastTriggerTs < c.IntervalNs {
		result.Triggered = false
		return result
	}

	result.Triggered = true
	c.LastTriggerTs = ts
	c.LastTriggerTotalCount = c.CurrentTotalCount
	return result
}
