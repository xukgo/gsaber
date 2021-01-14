package prioriyQueue

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/xukgo/gsaber/utils/randomUtil"
)

type Node struct {
	priority int
	value    int
}

func (this *Node) Less(other interface{}) bool {
	ov := other.(*Node)
	if this.priority < ov.priority {
		return true
	}
	if this.priority > ov.priority {
		return false
	}
	return this.value < ov.value
}

func TestPriorityQueue1(t *testing.T) {
	var count = 100000
	q := NewQueue(0)
	for i := 0; i < count; i++ {
		q.Enqueue(&Node{priority: int(randomUtil.NewInt32(1, 4)), value: int(randomUtil.NewInt32(5, 8))})
	}

	fmt.Println("queue len is", q.Len())
	sarr := make([]SortData, 0, count)
	for {
		item := q.Dequeue()
		if item == nil {
			break
		}
		sarr = append(sarr, item)
		//fmt.Println(item)
	}

	if !checkSortArrayValid(sarr) {
		t.FailNow()
	}
}

func TestPriorityQueue2(t *testing.T) {
	var count = 30
	q := NewQueue(0)

	var recvCnt int32 = 0
	for i := 0; i < 4; i++ {
		go func() {
			for {
				item := q.WaitDequeue()
				fmt.Println(item)
				atomic.AddInt32(&recvCnt, 1)
				if int(recvCnt) >= count {
					return
				}
			}
		}()
	}

	time.Sleep(time.Millisecond * 20)
	for i := 0; i < count; i++ {
		q.Enqueue(&Node{priority: int(randomUtil.NewInt32(1, 4)), value: int(randomUtil.NewInt32(5, 8))})
		time.Sleep(time.Millisecond * 1)
	}

	time.Sleep(time.Millisecond * 1000)
	//sarr := make([]SortData, 0, 32)
	//for {
	//	item := q.Dequeue()
	//	if item == nil {
	//		break
	//	}
	//	sarr = append(sarr, item)
	//	//fmt.Println(item)
	//}
	//
	//if !checkSortArrayValid(sarr) {
	//	t.FailNow()
	//}
}

func checkSortArrayValid(sarr []SortData) bool {
	for i := 0; i < len(sarr)-1; i++ {
		if !sarr[i].Less(sarr[i+1]) && sarr[i+1].Less(sarr[i]) {
			return false
		}
	}
	return true
}
