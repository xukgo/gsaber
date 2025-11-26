package periodTrigger

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_hashCollection_strKey(t *testing.T) {
	var target IPeriodTriggerCollection = NewHashCollection(time.Second*30, time.Second*3)

	var state TriggerState
	key1 := "abc_k1"
	state = target.Check(key1, 1000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	assert.True(t, state.IntervalNs == int64(time.Nanosecond*2000))
	assert.True(t, state.LastTriggerTs == 0)
	assert.True(t, state.LastTriggerTotalCount == 0)
	assert.True(t, state.CurrentTotalCount == 1)

	state = target.Check(key1, 1001, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.Check(key1, 1002, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	assert.True(t, state.IntervalNs == int64(time.Nanosecond*2000))
	assert.True(t, state.LastTriggerTs == 1000)
	assert.True(t, state.LastTriggerTotalCount == 1)
	assert.True(t, state.CurrentTotalCount == 3)

	state = target.Check(key1, 3000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)

	assert.True(t, target.Count() == 1)

	key2 := "abc_k2"
	state = target.Check(key2, 1000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	state = target.Check(key2, 1001, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.Check(key2, 1002, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.Check(key2, 4000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	assert.True(t, target.Count() == 2)

	var evictCount = target.Evict(1000)
	assert.True(t, evictCount == 0)
	assert.True(t, target.Count() == 2)

	evictCount = target.Evict(int64(time.Second*30) + int64(3000))
	assert.True(t, evictCount == 1)
	assert.True(t, target.Count() == 1)

	evictCount = target.Evict(int64(time.Second*30) + int64(4000))
	assert.True(t, evictCount == 1)
	assert.True(t, target.Count() == 0)
}

func Test_hashCollection_uintKey(t *testing.T) {
	var target IPeriodTriggerCollection = NewHashCollection(time.Second*30, time.Second*3)

	var state TriggerState
	var key1 uint64 = 1001
	state = target.CheckUint(key1, 1000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	assert.True(t, state.IntervalNs == int64(time.Nanosecond*2000))
	assert.True(t, state.LastTriggerTs == 0)
	assert.True(t, state.LastTriggerTotalCount == 0)
	assert.True(t, state.CurrentTotalCount == 1)

	state = target.CheckUint(key1, 1001, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.CheckUint(key1, 1002, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	assert.True(t, state.IntervalNs == int64(time.Nanosecond*2000))
	assert.True(t, state.LastTriggerTs == 1000)
	assert.True(t, state.LastTriggerTotalCount == 1)
	assert.True(t, state.CurrentTotalCount == 3)

	state = target.CheckUint(key1, 3000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)

	assert.True(t, target.Count() == 1)

	var key2 uint64 = 1002
	state = target.CheckUint(key2, 1000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	state = target.CheckUint(key2, 1001, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.CheckUint(key2, 1002, time.Nanosecond*2000)
	assert.True(t, state.Triggered == false)
	state = target.CheckUint(key2, 4000, time.Nanosecond*2000)
	assert.True(t, state.Triggered == true)
	assert.True(t, target.Count() == 2)

	var evictCount = target.Evict(1000)
	assert.True(t, evictCount == 0)
	assert.True(t, target.Count() == 2)

	evictCount = target.Evict(int64(time.Second*30) + int64(3000))
	assert.True(t, evictCount == 1)
	assert.True(t, target.Count() == 1)

	evictCount = target.Evict(int64(time.Second*30) + int64(4000))
	assert.True(t, evictCount == 1)
	assert.True(t, target.Count() == 0)
}

func Test_hashCollection02(t *testing.T) {
	var target IPeriodTriggerCollection = NewHashCollection(time.Second*30, time.Second*3)
	target.CheckDefaultFuncLine(1000, time.Second)
}
