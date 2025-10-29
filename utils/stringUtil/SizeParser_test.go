package stringUtil

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sizeText_unit_binary(t *testing.T) {
	target := NewSizeParser(true)
	res, err := target.Parse("1B")
	assert.True(t, err == nil && res == 1)
	res, err = target.Parse("9B")
	assert.True(t, err == nil && res == 9)
	res, err = target.Parse("9k")
	assert.True(t, err == nil && res == 9*1024)
	res, err = target.Parse("0.1k")
	assert.True(t, err == nil && res == 102)
	res, err = target.Parse("9M")
	assert.True(t, err == nil && res == 9*1024*1024)
	res, err = target.Parse("0.1M")
	assert.True(t, err == nil && res == int64(math.Floor(0.1*1024*1024)))
	res, err = target.Parse("9G")
	assert.True(t, err == nil && res == 9*1024*1024*1024)
	res, err = target.Parse("0.1G")
	assert.True(t, err == nil && res == int64(math.Floor(0.1*1024*1024*1024)))
}

func Test_sizeText_unit_hex(t *testing.T) {
	target := NewSizeParser(false)
	res, err := target.Parse("1B")
	assert.True(t, err == nil && res == 1)
	res, err = target.Parse("9B")
	assert.True(t, err == nil && res == 9)
	res, err = target.Parse("9k")
	assert.True(t, err == nil && res == 9*1000)
	res, err = target.Parse("0.1k")
	assert.True(t, err == nil && res == 100)
	res, err = target.Parse("9M")
	assert.True(t, err == nil && res == 9*1000*1000)
	res, err = target.Parse("0.1M")
	assert.True(t, err == nil && res == int64(math.Floor(0.1*1000*1000)))
	res, err = target.Parse("9G")
	assert.True(t, err == nil && res == 9*1000*1000*1000)
	res, err = target.Parse("0.1G")
	assert.True(t, err == nil && res == int64(math.Floor(0.1*1000*1000*1000)))
}
