package queueLocker

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

// Only items implementing this interface can be enqueued
// on the priority queue.
type SortData interface {
	Less(interface{}) bool
	Equal(interface{}) bool
}

//为优先级排队锁而创建的优先级队列，用的是链表不是堆解决方案，纯优先级队列在大数据的情况下不如另外那个，不一定适用于别的场景，使用要慎重
type priorityQueue struct {
	items    *list.List
	cond     *sync.Cond
	locker   *sync.Mutex
	waitData SortData
	minData  SortData
	//canCheckTop bool
}

// New creates and initializes a new priority queue, taking
// a limit as a parameter. If 0 given, then queue will be
// unlimited.
func newPriorityQueue(minData SortData) (q *priorityQueue) {
	q = &priorityQueue{}
	q.items = list.New()
	q.locker = new(sync.Mutex)
	q.cond = sync.NewCond(new(sync.Mutex))
	q.minData = minData
	//q.canCheckTop = true
	return
}

func (q *priorityQueue) addFrontMin() {
	q.items.PushFront(q.minData)
}

func (q *priorityQueue) removeFrontMin() {
	q.remove(q.minData)
}

// Enqueue puts given item to the queue.
func (q *priorityQueue) Enqueue(item SortData) {
	q.locker.Lock()
	q.insertSort(item)
	//fmt.Println("add item")
	q.locker.Unlock()

	if q.items.Len() == 1 {
		q.cond.Broadcast()
	} else {
		//q.cond.Signal()
	}

	return
}

func (q *priorityQueue) PopEqualTopWait(item SortData, timeout time.Duration) bool {
	startAt := time.Now()
	timeoutNs := timeout.Nanoseconds()

	for {
		q.locker.Lock()
		if q.checkTopEqual(item) {
			q.waitData = item
			q.addFrontMin()
			q.locker.Unlock()
			return true
		}

		if timeoutNs > 0 && time.Since(startAt).Nanoseconds() > timeoutNs {
			q.remove(item)
			q.locker.Unlock()
			return false
		}

		q.locker.Unlock()
		q.cond.L.Lock()
		//fmt.Println("pop cond wait")
		q.cond.Wait()
		q.cond.L.Unlock()
		continue
	}
}

func (q *priorityQueue) PopEqualTopRelease() {
	q.locker.Lock()
	q.remove(q.waitData)
	q.removeFrontMin()
	q.locker.Unlock()

	//fmt.Println("pop cond Broadcast")
	q.cond.Broadcast()
}

func (q *priorityQueue) Broadcast() {
	q.cond.Broadcast()
}

func (q *priorityQueue) remove(item SortData) {
	var s *list.Element = nil
	for p := q.items.Front(); p != nil; p = p.Next() {
		if p.Value.(SortData).Equal(item) {
			s = p
			break
		}
	}
	if s != nil {
		q.items.Remove(s)
	}
}

func (q *priorityQueue) Print() {
	q.locker.Lock()
	for p := q.items.Front(); p != nil; p = p.Next() {
		fmt.Printf("%v\n", p.Value)
	}
	q.locker.Unlock()
}

func (q *priorityQueue) insertSort(item SortData) {
	for p := q.items.Front(); p != nil; p = p.Next() {
		ov := p.Value.(SortData)
		if item.Less(ov) {
			q.items.InsertBefore(item, p)
			return
		}
	}

	q.items.PushBack(item)
}

func (q *priorityQueue) checkTopEqual(item SortData) bool {
	if q.items.Len() == 0 {
		return false
	}
	x := q.items.Front().Value.(SortData)
	br := x.Equal(item)
	return br
}
