package fileWatcher

import (
	"fmt"
	"sync"
	"time"
)

var localWatcherLocker = sync.RWMutex{}
var localWatchers = make([]*Watcher, 0, 0)

func Start() {
	fmt.Printf("start file watch service\r\n")
	minInterval := 200
	go func() {
		for {
			localWatcherLocker.RLock()
			for idx := range localWatchers {
				localWatchers[idx].do()
			}
			localWatcherLocker.RUnlock()
			time.Sleep(time.Duration(minInterval) * time.Millisecond)
		}
	}()
}

func AddWatcher(model *Watcher) {
	localWatcherLocker.Lock()
	localWatchers = append(localWatchers, model)
	localWatcherLocker.Unlock()
}
