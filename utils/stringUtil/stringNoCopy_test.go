package stringUtil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xukgo/gsaber/utils/randomUtil"
)

func Test_nocopy_string2Bytes(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, bytes.Equal(NoCopyString2Bytes(bf1.String()), NoCopyString2Bytes(bf2.String())))
	bf1.WriteString("a")
	bf2.WriteString("a")
	assert.True(t, bytes.Equal(NoCopyString2Bytes(bf1.String()), NoCopyString2Bytes(bf2.String())))
	bf1.WriteString("123")
	bf2.WriteString("123")
	assert.True(t, bytes.Equal(NoCopyString2Bytes(bf1.String()), NoCopyString2Bytes(bf2.String())))
	for i := 0; i < 1000; i++ {
		txt := randomUtil.NewString(randomUtil.NewInt(1, 10))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		assert.True(t, bytes.Equal(NoCopyString2Bytes(bf1.String()), NoCopyString2Bytes(bf2.String())))
	}
}

func Test_nocopy_bytes2String(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, NoCopyBytes2String(bf1.Bytes()) == NoCopyBytes2String(bf2.Bytes()))
	bf1.WriteString("a")
	bf2.WriteString("a")
	assert.True(t, NoCopyBytes2String(bf1.Bytes()) == NoCopyBytes2String(bf2.Bytes()))
	bf1.WriteString("123")
	bf2.WriteString("123")
	assert.True(t, NoCopyBytes2String(bf1.Bytes()) == NoCopyBytes2String(bf2.Bytes()))
	for i := 0; i < 1000; i++ {
		txt := randomUtil.NewString(randomUtil.NewInt(1, 10))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		assert.True(t, NoCopyBytes2String(bf1.Bytes()) == NoCopyBytes2String(bf2.Bytes()))
	}
}

func Test_nocopy_bytes2String_02(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	bf1.WriteString("123")
	bf2.WriteString("abc")
	assert.True(t, NoCopyBytes2String(bf1.Bytes()[:0]) == NoCopyBytes2String(bf2.Bytes())[:0])
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(1, 10))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		str1 := NoCopyBytes2String(bf1.Bytes()[:len(txt)])
		str2 := NoCopyBytes2String(bf2.Bytes()[:len(txt)])
		assert.True(t, str1 == str2)
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(11, 32))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		str1 := NoCopyBytes2String(bf1.Bytes()[:len(txt)])
		str2 := NoCopyBytes2String(bf2.Bytes()[:len(txt)])
		assert.True(t, str1 == str2)
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(33, 200))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		str1 := NoCopyBytes2String(bf1.Bytes()[:len(txt)])
		str2 := NoCopyBytes2String(bf2.Bytes()[:len(txt)])
		assert.True(t, str1 == str2)
	}
}
