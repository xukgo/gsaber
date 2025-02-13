package timeUtil

import (
	"runtime"
	"time"
)

func BusyWait(delay time.Duration) {
	end := time.Now().Add(delay)
	for time.Now().Before(end) {
		runtime.Gosched() // Avoid fully occupying the core
	}
}

func BusyWaitDeadline(end time.Time) {
	for time.Now().Before(end) {
		runtime.Gosched() // Avoid fully occupying the core
	}
}
