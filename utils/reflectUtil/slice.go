package reflectUtil

import "unsafe"

func UnsafeExtendBytes(old []byte, newSize int) []byte {
	return unsafe.Slice(unsafe.SliceData(old), newSize)
}
