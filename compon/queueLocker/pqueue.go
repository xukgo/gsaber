package queueLocker

import (
	"container/list"
	"fmt"
	"sync"
)

// Only items implementing this interface can be enqueued
// on the priority queue.
type SortData interface {
	Less(interface{}) bool
	Equal(interface{}) bool
}

//为优先级排队锁而创建的优先级队列，用的是链表不是堆解决方案，纯优先级队列在大数据的情况下不如另外那个，不一定适用于别的场景，使用要慎重
type priorityQueue struct {
	items  *list.List
	cond   *sync.Cond
	locker *sync.RWMutex
	//canCheckTop bool
}

// New creates and initializes a new priority queue, taking
// a limit as a parameter. If 0 given, then queue will be
// unlimited.
func newPriorityQueue() (q *priorityQueue) {
	q = &priorityQueue{}
	q.items = list.New()
	q.locker = new(sync.RWMutex)
	q.cond = sync.NewCond(new(sync.Mutex))
	//q.canCheckTop = true
	return
}

// Enqueue puts given item to the queue.
func (q *priorityQueue) Enqueue(item SortData) {
	q.locker.Lock()
	q.insertSort(item)
	//fmt.Println("add item")
	q.locker.Unlock()

	q.cond.L.Lock()
	if q.items.Len() == 1 {
		q.cond.Broadcast()
	} else {
		//q.cond.Signal()
	}
	q.cond.L.Unlock()

	return
}

func (q *priorityQueue) PopEqualTopWait(item SortData) SortData {
	q.cond.L.Lock()
	for {
		q.locker.Lock()
		if q.checkTopEqual(item) {
			x := q.items.Front()
			q.items.Remove(x)
			//fmt.Println(q.items.Len())
			q.locker.Unlock()

			item = x.Value.(SortData)
			//fmt.Println("pop done")
			return item
		}

		q.locker.Unlock()
		q.cond.Wait()
		continue
	}
}

func (q *priorityQueue) PopEqualTopRelease() {
	q.cond.Broadcast()
	q.cond.L.Unlock()
}

// Dequeue takes an item from the queue. If queue is empty
// then should block waiting for at least one item.
//func (q *priorityQueue) WaitDequeue() (item SortData) {
//	q.cond.L.Lock()
//start:
//	x := q.items.Front()
//	if x != nil {
//		q.items.Remove(x)
//	}
//	//x := heap.Pop(q.items)
//	if x == nil {
//		q.cond.Wait()
//		goto start
//	}
//	q.cond.L.Unlock()
//	item = x.Value.(SortData)
//	return
//}

func (q *priorityQueue) Print() {
	q.cond.L.Lock()
	for p := q.items.Front(); p != nil; p = p.Next() {
		fmt.Printf("%v\n", p.Value)
	}
	q.cond.L.Unlock()
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

// Safely changes enqueued items limit. When limit is set
// to 0, then queue is unlimited.
//func (q *priorityQueue) ChangeLimit(newLimit int) {
//	q.cond.L.Lock()
//	defer q.cond.L.Unlock()
//	q.Limit = newLimit
//}

// Len returns number of enqueued elemnents.
func (q *priorityQueue) Len() int {
	return q.items.Len()
}

// IsEmpty returns true if queue is empty.
func (q *priorityQueue) IsEmpty() bool {
	return q.Len() == 0
}
