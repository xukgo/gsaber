package redisUtils

import "github.com/xukgo/gsaber/utils/algorithUtils"

func CalcSlots(keys []string) []int {
	slots := make([]int, len(keys))
	for inx, key := range keys {
		slots[inx] = int(KeyHash(key))
	}
	return slots
}

func KeyHash(key string) uint16 {
	hashtag := ""
findHashTag:
	for i, s := range key {
		if s == '{' {
			for k := i; k < len(key); k++ {
				if key[k] == '}' {
					hashtag = key[i+1 : k]
					break findHashTag
				}
			}
		}
	}
	if len(hashtag) > 0 {
		key = hashtag
	}
	return algorithUtils.Crc16(key) & 0x3FFF
}
