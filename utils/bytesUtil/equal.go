package bytesUtil

import "github.com/xukgo/gsaber/utils/stringUtil"

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
