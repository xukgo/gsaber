package queueLocker

import (
	"time"

	"github.com/xukgo/gsaber/utils/randomUtil"
)

type Locker struct {
	pq *priorityQueue
}

func NewLocker() *Locker {
	model := new(Locker)
	model.pq = newPriorityQueue()
	return model
}

//priority,0优先级最高，数值越大优先级越低
func (this *Locker) LockPriority(priority int) {
	pv := initPValue(priority, time.Now().UnixNano(), randomUtil.NewLowerHexString(6))
	this.pq.Enqueue(&pv)
	this.pq.PopEqualTopWait(&pv)
}

func (this *Locker) Unlock() {
	this.pq.PopEqualTopRelease()
}

type pvalue struct {
	priority int
	ts       int64
	key      string
}

func initPValue(priority int, ts int64, key string) pvalue {
	return pvalue{
		priority: priority,
		ts:       ts,
		key:      key,
	}
}

func (this *pvalue) Less(other interface{}) bool {
	ov := other.(*pvalue)
	if this.priority < ov.priority {
		return true
	}
	if this.priority > ov.priority {
		return false
	}
	return this.ts < ov.ts
}

func (this *pvalue) Equal(other interface{}) bool {
	ov := other.(*pvalue)
	if this.priority != ov.priority {
		return false
	}
	if this.ts != ov.ts {
		return false
	}
	if this.key != ov.key {
		return false
	}
	return true
}
