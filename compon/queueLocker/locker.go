package queueLocker

import (
	"sync/atomic"
	"time"

	"github.com/xukgo/gsaber/utils/randomUtil"
)

type Locker struct {
	pq        *priorityQueue
	interval  time.Duration
	tm        *time.Timer
	tmEnable  int32
	waitCount int32
	nextDoTs  int64
}

func NewLocker(interval time.Duration) *Locker {
	if interval == 0 {
		interval = time.Millisecond * 200
	}
	model := new(Locker)
	minData := initPValue(-0xffffffff, 0, "")
	model.pq = newPriorityQueue(&minData)
	model.interval = interval
	model.tmEnable = 0
	model.waitCount = 0
	return model
}

func (this *Locker) Dispose() {
	atomic.StoreInt32(&this.tmEnable, 0)
	this.tm = nil
}

func (this *Locker) startTimer() {
	if this.interval.Nanoseconds() == 0 {
		return
	}
	this.tm = time.NewTimer(this.interval)
	tm := this.tm
	go func() {
		for {
			<-tm.C
			//fmt.Println("timer")
			if atomic.LoadInt32(&this.tmEnable) == 0 {
				return
			}
			if time.Now().UnixNano() >= this.nextDoTs && this.waitCount > 0 {
				this.pq.Broadcast()
			}
			tm.Reset(this.interval)
		}
	}()

}

//priority,0优先级最高，数值越大优先级越低,timeout为0表示一直等下去
func (this *Locker) LockPriority(priority int, timeout time.Duration) bool {
	return this.LockPriorityTime(priority, time.Now(), timeout)
}

func (this *Locker) LockPriorityTime(priority int, dt time.Time, timeout time.Duration) bool {
	//懒汉模式，第一个锁操作才启动timer
	if atomic.CompareAndSwapInt32(&this.tmEnable, 0, 1) {
		this.startTimer()
	}

	pv := initPValue(priority, dt.UnixNano(), randomUtil.NewLowerHexString(6))
	this.pq.Enqueue(&pv)
	atomic.AddInt32(&this.waitCount, 1)
	br := this.pq.PopEqualTopWait(&pv, timeout)
	atomic.AddInt32(&this.waitCount, -1)
	return br
}

func (this *Locker) Unlock() {
	this.pq.PopEqualTopRelease()
	if this.interval.Nanoseconds() > 0 {
		this.nextDoTs = time.Now().Add(this.interval).UnixNano()
	}
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
