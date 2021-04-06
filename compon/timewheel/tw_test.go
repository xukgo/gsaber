package timewheel

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_timewheel(t *testing.T) {
	interval := time.Millisecond * 100
	tw := New(interval, 6, func(key interface{}, data interface{}) {
		fmt.Printf("%s action %v\n", formatTime(time.Now()), data)
	})
	time.Sleep(time.Second * 1)
	tw.Start()
	//tw.Add(time.Millisecond*100,"k1","d1",false)
	//tw.Remove("k1")
	//tw.Add(time.Millisecond*200,"k1","d2",false)
	//tw.Remove("k1")
	tw.AddCron(interval*2, "k1", "d3")
	//time.Sleep(time.Second*1)
	//tw.Remove("k1")
	time.Sleep(time.Second * 10)
}

func Test_circlePos(t *testing.T) {
	v := int(1 / 3)
	if v != 0 {
		t.FailNow()
	}
	v = int(2 / 3)
	if v != 0 {
		t.FailNow()
	}

	var pos, circle int
	interval := time.Duration(1)
	slotNum := 3
	pos, circle = getPositionAndCircle(time.Duration(0), interval, slotNum, 0)
	if pos != 0 || circle != 0 {
		t.FailNow()
	}
	pos, circle = getPositionAndCircle(time.Duration(1), interval, slotNum, 0)
	if pos != 1 || circle != 0 {
		t.FailNow()
	}
	pos, circle = getPositionAndCircle(time.Duration(2), interval, slotNum, 0)
	if pos != 2 || circle != 0 {
		t.FailNow()
	}
	pos, circle = getPositionAndCircle(time.Duration(3), interval, slotNum, 0)
	if pos != 0 || circle != 1 {
		t.FailNow()
	}

	pos, circle = getPositionAndCircle(time.Duration(6), interval, slotNum, 0)
	if pos != 0 || circle != 2 {
		t.FailNow()
	}
	pos, circle = getPositionAndCircle(time.Duration(5), interval, slotNum, 1)
	if pos != 0 || circle != 1 {
		t.FailNow()
	}
	pos, circle = getPositionAndCircle(time.Duration(9), 1, 5, 3)
	if pos != 2 || circle != 1 {
		t.FailNow()
	}

	pos, circle = getPositionAndCircle(time.Duration(5), 1, 6, 5)
	if pos != 4 || circle != 0 {
		t.FailNow()
	}
}

func TestSyncAddDelete(t *testing.T) {
	dict := new(sync.Map)
	k := "id1"
	v := "v1"
	dict.Store(k, v)
	dict.Range(func(key, value interface{}) bool {
		fmt.Printf("key=%v val=%v\n", key, value)
		dict.Store(key, value)
		dict.Delete(key)
		dict.Store(key, value)
		return true
	})

	dict.Range(func(key, value interface{}) bool {
		fmt.Printf("key2=%v val2=%v\n", key, value)
		return true
	})
}
