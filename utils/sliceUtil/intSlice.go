package sliceUtil

import "github.com/xukgo/gsaber/utils/randomUtil"

func RemoveMatchInt(arr []int, match int) []int {
	j := 0
	for idx := range arr {
		if arr[idx] != match {
			arr[j] = arr[idx]
			j++
		}
	}
	return arr[:j]
}

func RandomIntSliceIndex(arr []int) {
	if len(arr) <= 0 || len(arr) == 1 {
		return
	}

	for i := len(arr) - 1; i > 0; i-- {
		num := randomUtil.NewInt32(0, int32(i+1))
		arr[i], arr[num] = arr[num], arr[i]
	}
}
