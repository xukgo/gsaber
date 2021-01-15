package queueLocker

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/xukgo/gsaber/utils/randomUtil"
)

func TestQueueLocker(t *testing.T) {
	wg := sync.WaitGroup{}
	qlocker := NewLocker()
	var index int32 = 0
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			idx := atomic.AddInt32(&index, 1)
			//time.Sleep(time.Millisecond * time.Duration(20))
			//startAt := time.Now()
			priority := int(randomUtil.NewInt32(0, 3))
			qlocker.LockPriority(priority, 0)
			//fmt.Printf("wait %d ms, index %d\n", time.Since(startAt).Milliseconds(), idx)
			doQueueLockerAction(idx, priority)
			qlocker.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}

var do = 1

func doQueueLockerAction(index int32, priority int) {
	if do == 1 {
		time.Sleep(time.Millisecond * 30)
		do++
	}
	fmt.Printf("action %d %d\n", priority, index)
}

func TestQueueLockerLongTime(t *testing.T) {
	wg := sync.WaitGroup{}
	qlocker := NewLocker()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		dur := time.Millisecond //* time.Duration(randomUtil.NewInt32(10, 50))
		time.AfterFunc(dur, func() {
			priority := int(randomUtil.NewInt32(0, 3))
			br := qlocker.LockPriority(priority, time.Millisecond*20)
			if !br {
				//fmt.Println("lock failed")
			} else {
				//time.Sleep(time.Millisecond * time.Duration(randomUtil.NewInt32(10, 100)))
				fmt.Println("lock ok")
				qlocker.Unlock()
			}
			wg.Done()
		})
	}
	wg.Wait()
}
