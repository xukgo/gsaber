package bytesUtil

import (
	"math/bits"
	"unsafe"
)

// BytesFindNewLineIn8Bytes 在8字节块中查找换行符
// 返回值: 找到则返回字节索引(0-7)，未找到返回-1
func BytesFindNewLineIn8Bytes(data []byte) int {
	if len(data) < 8 {
		return -1
	}

	val := *(*uint64)(unsafe.Pointer(&data[0]))

	// SWAR算法查找 '\n' (0x0A)
	xor := val ^ 0x0A0A0A0A0A0A0A0A
	hasMatch := (xor - 0x0101010101010101) & ^xor & 0x8080808080808080

	if hasMatch != 0 {
		return bits.TrailingZeros64(hasMatch) / 8
	}
	return -1
}

// BytesFindNewLineIn4Bytes 在4字节块中查找换行符
// 返回值: 找到则返回字节索引(0-3)，未找到返回-1
func BytesFindNewLineIn4Bytes(data []byte) int {
	if len(data) < 4 {
		return -1
	}

	val := *(*uint32)(unsafe.Pointer(&data[0]))

	// SWAR算法查找 '\n' (0x0A)
	xor := val ^ 0x0A0A0A0A
	hasMatch := (xor - 0x01010101) & ^xor & 0x80808080

	if hasMatch != 0 {
		return bits.TrailingZeros32(hasMatch) / 8
	}
	return -1
}
