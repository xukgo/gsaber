package redisUtils

import (
	"bytes"

	"github.com/xukgo/gsaber/utils/stringUtil"
)

var charArr = "0123456789abcdefghjiklmnopqrstuvwxyz"

func GenerateStaticKey(slot int, len int) string {
	buff := bytes.Repeat([]byte{'0'}, len)
	for {
		if KeyHash(stringUtil.NoCopyBytes2String(buff)) == uint16(slot) {
			return string(buff)
		}
		buffIncreaseVal(buff)
	}
}

func buffIncreaseVal(buff []byte) {
	index := len(buff) - 1
	for {
		if index < 0 {
			return
		}
		if buff[index] == 'z' {
			buff[index] = '0'
			index--
			continue
		}
		if buff[index] == '9' {
			buff[index] = 'a'
			return
		}
		buff[index]++
		return
	}
}
