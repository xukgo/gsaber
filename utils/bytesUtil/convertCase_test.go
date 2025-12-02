package bytesUtil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toUpper(t *testing.T) {
	var data []byte
	data = []byte("")
	ToUpperAsciiSelf(data)
	assert.True(t, string(data) == "")

	data = []byte("123")
	ToUpperAsciiSelf(data)
	assert.True(t, string(data) == "123")

	data = []byte("abc123")
	ToUpperAsciiSelf(data)
	assert.True(t, string(data) == "ABC123")

	data = []byte("ABC123")
	ToUpperAsciiSelf(data)
	assert.True(t, string(data) == "ABC123")

}

func Test_toLower(t *testing.T) {
	var data []byte
	data = []byte("")
	ToLowerAsciiSelf(data)
	assert.True(t, string(data) == "")

	data = []byte("123")
	ToLowerAsciiSelf(data)
	assert.True(t, string(data) == "123")

	data = []byte("ABC123")
	ToLowerAsciiSelf(data)
	assert.True(t, string(data) == "abc123")

	data = []byte("abc123")
	ToLowerAsciiSelf(data)
	assert.True(t, string(data) == "abc123")
}
