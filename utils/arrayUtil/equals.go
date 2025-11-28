package arrayUtil

import "github.com/xukgo/gsaber/utils/stringUtil"

// EqualsInt returns a bool value indicating whether the int64[] arr1 and arr2 are equal
func EqualsInt(arr1, arr2 []int64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}
	arr2 = arr2[:len(arr1)]
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

// EqualsString returns a bool value indicating whether the string[] arr1 and arr2 are equal
func EqualsString(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	if (arr1 == nil) != (arr2 == nil) {
		return false
	}
	arr2 = arr2[:len(arr1)]
	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}
	return true
}

// EqualsBytes returns a bool value indicating whether the byte[] arr1 and arr2 are equal
func EqualsBytes(arr1, arr2 []byte) bool {
	if len(arr1) == 0 && len(arr2) == 0 {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}

func EqualsStringBytes(arr1, arr2 []byte) bool {
	if len(arr1) == 0 && len(arr2) == 0 {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}

	//str1 := stringUtil.NoCopyBytes2String(arr1)
	//str2 := stringUtil.NoCopyBytes2String(arr2)
	//return str1 == str2
	return stringUtil.NoCopyBytes2String(arr1) == stringUtil.NoCopyBytes2String(arr2)
}
