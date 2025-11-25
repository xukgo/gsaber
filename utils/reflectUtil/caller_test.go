package reflectUtil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetCallerLayerInfo(t *testing.T) {
	info := GetCallerLayerInfo(0)
	assert.True(t, info.Ok)
	assert.True(t, strings.HasSuffix(info.FileName, "test.go"))
	assert.True(t, strings.HasPrefix(info.Function, "Test_"))
}

func Test_GetCallerLayerInfoFromPC(t *testing.T) {
	pc := GetCallerPC(0)
	info := GetCallerLayerInfoFromPC(pc)
	assert.True(t, info.Ok)
	assert.True(t, strings.HasSuffix(info.FileName, "test.go"))
	assert.True(t, strings.HasPrefix(info.Function, "Test_"))
}

func Benchmark_GetCallerLayerInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetCallerLayerInfo(0)
	}
}

func Benchmark_GetCallerPC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc := GetCallerPC(0)
		GetCallerLayerInfoFromPC(pc)
	}
}
