package hashLocker

import (
	"github.com/spaolacci/murmur3"
	"sync"
)

type HashLocker struct {
	lockers []sync.Locker
	size    int
}

func NewHashLocker() *HashLocker {
	return NewSizeHashLocker(1024)
}
func NewSizeHashLocker(size int) *HashLocker {
	model := new(HashLocker)
	model.size = size
	model.lockers = make([]sync.Locker, size)
	for i := 0; i < size; i++ {
		model.lockers[i] = new(sync.Mutex)
	}
	return model
}

func (this *HashLocker) Lock(key []byte) {
	lk := this.lockers[this.getKeyIndex(key)]
	lk.Lock()
}

func (this *HashLocker) Unlock(key []byte) {
	lk := this.lockers[this.getKeyIndex(key)]
	lk.Unlock()
}

func (this *HashLocker) getKeyIndex(key []byte) int {
	sum := murmur3.Sum32(key)
	return (int(sum)) % this.size
}
