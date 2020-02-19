package sliceUtil

import "strings"

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
