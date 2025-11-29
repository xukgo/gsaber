package arrayUtil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindNewLineIn4Bytes(t *testing.T) {
	var pos int
	pos = BytesFindNewLineIn8Bytes([]byte("12345678"))
	assert.True(t, pos == -1)
	pos = BytesFindNewLineIn8Bytes([]byte("\n2345678"))
	assert.True(t, pos == 0)
	pos = BytesFindNewLineIn8Bytes([]byte("1\n2345678"))
	assert.True(t, pos == 1)
	pos = BytesFindNewLineIn8Bytes([]byte("12\n345678"))
	assert.True(t, pos == 2)
	pos = BytesFindNewLineIn8Bytes([]byte("123\n45678"))
	assert.True(t, pos == 3)
	pos = BytesFindNewLineIn8Bytes([]byte("1234\n5678"))
	assert.True(t, pos == 4)
	pos = BytesFindNewLineIn8Bytes([]byte("12345\n678"))
	assert.True(t, pos == 5)
	pos = BytesFindNewLineIn8Bytes([]byte("123456\n78"))
	assert.True(t, pos == 6)
	pos = BytesFindNewLineIn8Bytes([]byte("1234567\n"))
	assert.True(t, pos == 7)

	pos = BytesFindNewLineIn4Bytes([]byte("1234"))
	assert.True(t, pos == -1)
	pos = BytesFindNewLineIn4Bytes([]byte("\n123"))
	assert.True(t, pos == 0)
	pos = BytesFindNewLineIn4Bytes([]byte("1\n23"))
	assert.True(t, pos == 1)
	pos = BytesFindNewLineIn4Bytes([]byte("12\n3"))
	assert.True(t, pos == 2)
	pos = BytesFindNewLineIn4Bytes([]byte("123\n"))
	assert.True(t, pos == 3)
}

func Benchmark_FindNewLineIn4Bytes_01(b *testing.B) {
	data := []byte("123\n")
	for i := 0; i < b.N; i++ {
		BytesFindNewLineIn4Bytes(data)
	}
}

func Benchmark_BytesFind4Bytes_01(b *testing.B) {
	data := []byte("123\n")
	for i := 0; i < b.N; i++ {
		bytes.IndexByte(data, '\n')
	}
}

func Benchmark_FindNewLineIn8Bytes_01(b *testing.B) {
	data := []byte("1234567\n")
	for i := 0; i < b.N; i++ {
		BytesFindNewLineIn8Bytes(data)
	}
}

func Benchmark_BytesFind8Bytes_01(b *testing.B) {
	data := []byte("1234567\n")
	for i := 0; i < b.N; i++ {
		bytes.IndexByte(data, '\n')
	}
}
