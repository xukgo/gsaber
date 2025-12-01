package stringUtil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toUpper(t *testing.T) {
	sb := &strings.Builder{}
	sb.Grow(12)

	ToUpper("", sb)
	assert.True(t, sb.String() == "")
	sb.Reset()
	ToUpper("123", sb)
	assert.True(t, sb.String() == "123")
	sb.Reset()
	ToUpper("abc123", sb)
	assert.True(t, sb.String() == "ABC123")
	sb.Reset()
	ToUpper("ABC123", sb)
	assert.True(t, sb.String() == "ABC123")
}

func Test_toLower(t *testing.T) {
	sb := &strings.Builder{}
	sb.Grow(12)

	ToLower("", sb)
	assert.True(t, sb.String() == "")
	sb.Reset()
	ToLower("123", sb)
	assert.True(t, sb.String() == "123")
	sb.Reset()
	ToLower("abc123", sb)
	assert.True(t, sb.String() == "abc123")
	sb.Reset()
	ToLower("ABC123", sb)
	assert.True(t, sb.String() == "abc123")
}
