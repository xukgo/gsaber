package arrayUtil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xukgo/gsaber/utils/randomUtil"
)

func Test_EqualsString(t *testing.T) {
	var arr1 = []string{"wxnacy", "wen", "go"}
	var arr2 = []string{"wxnacy", "wen", "go"}
	flag := EqualsString(arr1, arr2)
	if !flag {
		t.Errorf("%v is error", flag)
	}
	arr1 = []string{"wxnacy", "go"}
	arr2 = []string{"wxnacy", "wen", "go"}
	flag = EqualsString(arr1, arr2)
	if flag {
		t.Errorf("%v is error", flag)
	}
}

func Test_EqualsInt(t *testing.T) {
	var arr1 = []int64{1, 2, 3}
	var arr2 = []int64{1, 2, 3}
	flag := EqualsInt(arr1, arr2)
	if !flag {
		t.Errorf("%v is error", flag)
	}
	arr1 = []int64{1, 2}
	arr2 = []int64{3, 4}
	flag = EqualsInt(arr1, arr2)
	if flag {
		t.Errorf("%v is error", flag)
	}
}

func Test_equalBytes_mustSuccess_01(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(1, 10))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(11, 32))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(33, 200))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
	}
}
func Test_equalBytes_mustSuccess_02(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
	bf1.WriteString("123")
	bf2.WriteString("abc")
	assert.True(t, EqualsBytes(bf1.Bytes()[:0], bf2.Bytes()[:0]))
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(1, 10))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		assert.True(t, EqualsBytes(bf1.Bytes()[:len(txt)], bf2.Bytes()[:len(txt)]))
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(11, 32))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		assert.True(t, EqualsBytes(bf1.Bytes()[:len(txt)], bf2.Bytes()[:len(txt)]))
	}
	for i := 0; i < 1000; i++ {
		bf1.Reset()
		bf2.Reset()
		txt := randomUtil.NewString(randomUtil.NewInt(33, 200))
		bf1.WriteString(txt)
		bf2.WriteString(txt)
		bf1.WriteString("123")
		bf2.WriteString("abc")
		assert.True(t, EqualsBytes(bf1.Bytes()[:len(txt)], bf2.Bytes()[:len(txt)]))
	}
}

func Test_equalBytes_mustFail_01(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("1")
	bf2.WriteString("2")
	assert.False(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("111")
	bf2.WriteString("222")
	assert.False(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("122")
	bf2.WriteString("123")
	assert.False(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("123")
	bf1.WriteByte(0x0)
	bf1.WriteString("1")

	bf2.WriteString("123")
	bf2.WriteByte(0x0)
	bf2.WriteString("2")
	assert.False(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("123")
	bf1.WriteByte(0x0)
	bf1.WriteString("a")
	bf1.WriteString("b")

	bf2.WriteString("123")
	bf2.WriteString("a")
	bf2.WriteByte(0x0)
	bf2.WriteString("b")
	assert.False(t, EqualsBytes(bf1.Bytes(), bf2.Bytes()))
}

func Test_equalStringBytes_mustFail_01(t *testing.T) {
	bf1 := bytes.NewBuffer(nil)
	bf2 := bytes.NewBuffer(nil)
	assert.True(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("1")
	bf2.WriteString("2")
	assert.False(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("111")
	bf2.WriteString("222")
	assert.False(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("122")
	bf2.WriteString("123")
	assert.False(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("123")
	bf1.WriteByte(0x0)
	bf1.WriteString("1")

	bf2.WriteString("123")
	bf2.WriteByte(0x0)
	bf2.WriteString("2")
	assert.False(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))

	bf1.Reset()
	bf2.Reset()
	bf1.WriteString("123")
	bf1.WriteByte(0x0)
	bf1.WriteString("a")
	bf1.WriteString("b")

	bf2.WriteString("123")
	bf2.WriteString("a")
	bf2.WriteByte(0x0)
	bf2.WriteString("b")
	assert.False(t, EqualsStringBytes(bf1.Bytes(), bf2.Bytes()))
}
