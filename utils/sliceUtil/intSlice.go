package sliceUtil

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
