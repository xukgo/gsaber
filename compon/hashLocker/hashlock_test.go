package hashLocker

import "testing"

func TestHashLocker(t *testing.T) {
	hashLockerMap := NewLocker()

	key := []byte("gjfgkhgfkuehkht87343756396539785643htgeg4454;'[;]")

	for i := 0; i < 1000000; i++ {
		hashLockerMap.Lock(key)
		hashLockerMap.Unlock(key)
	}
}
