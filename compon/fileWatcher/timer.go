package fileWatcher

import (
	"sync"
	"time"
)

var localWatcherLocker = sync.RWMutex{}
var localWatchers = make([]*Watcher, 0, 0)

func Start() {
	//fmt.Printf("start file watch service\r\n")
	update()

	minInterval := 200
	go func() {
		for {
			time.Sleep(time.Duration(minInterval) * time.Millisecond)
			update()
		}
	}()
}

func AddWatcher(model *Watcher) {
	localWatcherLocker.Lock()
	localWatchers = append(localWatchers, model)
	localWatcherLocker.Unlock()
}

func update() {
	localWatcherLocker.RLock()
	for idx := range localWatchers {
		localWatchers[idx].do()
	}
	localWatcherLocker.RUnlock()
}
