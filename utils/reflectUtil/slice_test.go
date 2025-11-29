package reflectUtil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_unsafeExtendBytes(t *testing.T) {
	raw := []byte("0123456789abcedfg")
	s1 := raw[:5]
	s2 := s1[:3]
	s3 := UnsafeExtendBytes(s1, 10)
	s4 := UnsafeExtendBytes(s2, 10)
	assert.True(t, string(s3) == "0123456789")
	assert.True(t, string(s4) == "0123456789")
}
