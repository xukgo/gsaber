package randomUtil

import (
	"crypto/rand"
	"encoding/binary"
	"strconv"
	"strings"
)

var charArr1 =  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var charArr2 =  "0123456789ABCDEF"

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func newInt32()int32{
	bs := randBytes(4)
	longVal := int32(binary.BigEndian.Uint32(bs))
	return longVal
}

func newInt64()int64{
	bs := randBytes(8)
	longVal := int64(binary.BigEndian.Uint64(bs))
	return longVal
}

//x>=min and x<max
func NewInt32(min, max int32)int32{
	randVal := newInt32()
	if randVal < 0{
		randVal = -randVal
	}

	seg := int32(max-min)
	randVal = randVal%seg
	return randVal+min
}

func NewInt64(min, max int64)int64{
	randVal := newInt64()
	if randVal < 0{
		randVal = -randVal
	}
	seg := max-min
	randVal = randVal%seg
	return randVal+min
}

func NewHexString(count int)string{
	var bb []byte
	letters := []byte(charArr2)
	runeLen := int32(len(letters))
	for i := 0; i < count; i++ {
		ru := letters[NewInt32(0,runeLen)]
		bb = append(bb, ru)
	}
	return string(bb)
}

//新建一个随机手机号，1开头的11位，后面不管
func NewPhoneNumber()string{
	var builder strings.Builder

	builder.WriteString("1")
	for i:=0;i<10;i++{
		builder.WriteString(strconv.Itoa(int(NewInt32(0,10))))
	}

	return builder.String()
}