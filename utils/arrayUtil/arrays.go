package arrayUtil

import "strings"

func StringsTrimSpaceFilterEmpty(arr []string) []string {
	res := make([]string, 0, len(arr))
	for _, item := range arr {
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}
		res = append(res, item)
	}
	return res
}

func StringsToLower(arr []string) []string {
	res := make([]string, 0, len(arr))
	for _, item := range arr {
		res = append(res, strings.ToLower(item))
	}
	return res
}

func StringsToUpper(arr []string) []string {
	res := make([]string, 0, len(arr))
	for _, item := range arr {
		res = append(res, strings.ToUpper(item))
	}
	return res
}

func StringsContains(array []string, val string) (index int) {
	index = ContainsString(array, val)
	return
}

func StringsContainsAll(array []string, vals []string) bool {
	matchCount := 0
	for idx := range vals {
		if ContainsString(array, vals[idx]) >= 0 {
			matchCount++
		}
	}

	if matchCount == len(vals) {
		return true
	}
	return false
}

// func IntsContains(array []int, val int) (index int) {
// index = ContainsInt(array, val)
// return
// }

// func FloatsContains(array []float64, val float64) (index int) {
// index = ContainsFloat64(array, val)
// return
// }

// []string deduplicate
func StringsDeduplicate(array []string) []string {
	var arr = make([]string, 0)
	var m = make(map[string]bool)
	for _, d := range array {
		_, ok := m[d]
		if !ok {
			m[d] = true
			arr = append(arr, d)
		}
	}
	return arr
}

// []string equal
func StringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// []int deduplicate
func IntsDeduplicate(array []int) []int {
	var arr = make([]int, 0)
	var m = make(map[int]bool)
	for _, d := range array {
		_, ok := m[d]
		if !ok {
			m[d] = true
			arr = append(arr, d)
		}
	}
	return arr
}

// []int equal
func IntsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
