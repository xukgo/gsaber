package hashLocker

import (
	"github.com/spaolacci/murmur3"
	"sync"
)

type Locker struct {
	lockers []sync.Locker
	size    int
}

func NewLocker() *Locker {
	return NewSizeLocker(1024)
}
func NewSizeLocker(size int) *Locker {
	model := new(Locker)
	model.size = size
	model.lockers = make([]sync.Locker, size)
	for i := 0; i < size; i++ {
		model.lockers[i] = new(sync.Mutex)
	}
	return model
}

func (this *Locker) Lock(key []byte) int{
	index := this.GetKeyIndex(key)
	lk := this.lockers[index]
	lk.Lock()
	return index
}

func (this *Locker) Unlock(key []byte) int{
	index := this.GetKeyIndex(key)
	lk := this.lockers[index]
	lk.Unlock()
	return index
}

func (this *Locker) GetKeyIndex(key []byte) int {
	sum := murmur3.Sum32(key)
	return (int(sum)) % this.size
}

func (this *Locker) GetKeyHash(key []byte) uint32 {
	sum := murmur3.Sum32(key)
	return sum
}
