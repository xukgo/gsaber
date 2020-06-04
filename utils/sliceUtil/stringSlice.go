package sliceUtil

import (
	"github.com/xukgo/gsaber/utils/randomUtil"
	"strings"
)

func RemoveMatchString(arr []string, match string, matchCase bool) []string {
	j := 0
	for idx := range arr {
		if matchCase {
			if arr[idx] != match {
				arr[j] = arr[idx]
				j++
			}
		} else {
			if !strings.EqualFold(arr[idx], match) {
				arr[j] = arr[idx]
				j++
			}
		}
	}
	return arr[:j]
}

func RandomStringSliceIndex(arr []int) {
	if len(arr) <= 0 || len(arr) == 1 {
		return
	}

	for i := len(arr) - 1; i > 0; i-- {
		num := randomUtil.NewInt32(0, int32(i+1))
		arr[i], arr[num] = arr[num], arr[i]
	}
}
