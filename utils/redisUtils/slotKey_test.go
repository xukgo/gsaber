package redisUtils

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xukgo/gsaber/utils/stringUtil"
)

func Test_buffIncreaseVal(t *testing.T) {
	buff := bytes.Repeat([]byte{'0'}, 8)

	for i := 0; i < 100; i++ {
		fmt.Println(stringUtil.NoCopyBytes2String(buff))
		buffIncreaseVal(buff)
	}
}

func Test_GenerateStaticKey(t *testing.T) {
	for i := 0; i < 0x3FFF; i++ {
		txt := GenerateStaticKey(i, 4)
		key := fmt.Sprintf("{%s}", txt)
		slot := KeyHash(key)
		assert.True(t, i == int(slot))
		//fmt.Printf("[%d]-> %s\n", i, txt)
	}
}
