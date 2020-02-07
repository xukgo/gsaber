package randomUtil

import (
	"encoding/binary"
	"github.com/valyala/fastrand"
	"strconv"
	"strings"
)

var charArr1 =  "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var charArr2 =  []byte("0123456789ABCDEF")

func rand8Bytes() []byte {
	var buf = make([]byte, 8)
	mb := buf[4:]
	binary.BigEndian.PutUint32(buf, fastrand.Uint32())
	binary.BigEndian.PutUint32(mb, fastrand.Uint32())
	return buf
}


func newInt32()int32{
	n := fastrand.Uint32()
	return int32(n)
}

func newInt64()int64{
	bs := rand8Bytes()
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
	var bb = make([]byte,count)
	letters := charArr2
	runeLen := int32(len(letters))
	for i := 0; i < count; i++ {
		ru := letters[NewInt32(0,runeLen)]
		bb[i] = ru
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