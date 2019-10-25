package fileWatcher

import (
	"fmt"
	"github.com/xukgo/gsaber/utils/fileUtil"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Watcher struct {
	interval       int
	fileUrl        string
	lastWriteTime  int64
	observerArray  []Observer
	observerLocker sync.Mutex
}

func NewWatcher(interval int, url string) *Watcher {
	model := &Watcher{
		interval:      interval,
		fileUrl:       fileUtil.GetAbsUrl(url),
		lastWriteTime: 0,
		observerArray: []Observer{},
	}
	return model
}

func (this *Watcher) AddObserver(observer Observer) {
	this.observerLocker.Lock()
	this.observerArray = append(this.observerArray, observer)
	this.observerLocker.Unlock()
}

func (this *Watcher) RemoveObserver(observer Observer) {
	this.observerLocker.Lock()
	defer this.observerLocker.Unlock()

	index := -1
	for i := 0; i < len(this.observerArray); i++ {
		if this.observerArray[i] == observer {
			index = 1
			break
		}
	}

	if index < 0 {
		return
	}

	this.observerArray = append(this.observerArray[:index], this.observerArray[index+1:]...)
}

func (this *Watcher) Start() {
	fmt.Println("start file watch url:" + this.fileUrl)
	this.do()

	go func(instance *Watcher) {
		for {
			instance.do()
			time.Sleep(time.Duration(instance.interval) * time.Second)
		}
	}(this)
}

func (this *Watcher) do() {
	fileInfo, err := os.Stat(this.fileUrl)
	if err != nil {
		fmt.Printf("file watcher stat file return error; url:%s, err:%s", this.fileUrl, err.Error())
		return
	}

	modTs := fileInfo.ModTime().Unix()
	if modTs != this.lastWriteTime {
		fileContent, err := ioutil.ReadFile(this.fileUrl)
		if err != nil {
			fmt.Printf("file watcher ReadFile return error; url:%s, err:%s", this.fileUrl, err.Error())
			return
		}

		this.observerLocker.Lock()
		if this.observerArray != nil && len(this.observerArray) > 0 {
			fmt.Printf("file changed notify observers; url:%s", this.fileUrl)
			for _, observer := range this.observerArray {
				observer.UpdateFromContent(fileContent)
			}
		}
		this.observerLocker.Unlock()

		this.lastWriteTime = modTs
	}
}
