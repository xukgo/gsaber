package stringUtil

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unsafe"

	"github.com/xukgo/gsaber/utils/ruleUtil"
)

// 判断src偏移位置开始是否以match开头,不做空字符串的匹配
func StartWithIndex(src string, srcIndex int, sub string) bool {
	subLen := len(sub)
	if subLen == 0 {
		return false
	}
	srclen := len(src) - srcIndex
	if srclen == 0 {
		return false
	}
	if srclen < subLen {
		return false
	}

	for i := 0; i < subLen; i++ {
		if src[i+srcIndex] != sub[i] {
			return false
		}
	}

	return true
}

// 判断src是否以match开头,不做空字符串的匹配
func StartWith(src string, sub string) bool {
	subLen := len(sub)
	if subLen == 0 {
		return false
	}
	srclen := len(src)
	if srclen == 0 {
		return false
	}
	if srclen < subLen {
		return false
	}

	for i := 0; i < subLen; i++ {
		if src[i] != sub[i] {
			return false
		}
	}

	return true
}

// 按照分隔符分隔，并且分隔符也作为独立的字符串纳入返回结果
func SplitContainsSeps(src string, seps []string) []string {
	var arr []string
	srclen := len(src)

	index := 0 //代表src指针目前的偏移
	bf := new(bytes.Buffer)
	for i := 0; i < srclen; i++ {
		sepLen := 0
		matchSeq := ""
		for idx := range seps {
			if StartWithIndex(src, index, seps[idx]) {
				matchSeq = seps[idx]
				sepLen = len(matchSeq)
				break
			}
		}
		if sepLen == 0 {
			bf.WriteByte(src[i])
			index++
		} else {
			if bf.Len() > 0 {
				arr = append(arr, bf.String())
				bf.Reset()
			}
			arr = append(arr, matchSeq)
			index += sepLen
		}
	}

	//结尾后处理，如果是数字还要加进去
	if bf.Len() > 0 {
		arr = append(arr, bf.String())
		bf.Reset()
	}
	return arr
}

// 删除号码前置的86，区号要补0
func TelNoDelete86Head(number string) string {
	if strings.Index(number, "86") != 0 {
		return number
	}

	number = number[2:]

	if ruleUtil.CheckIsCnMobil(number) {
		return number
	}

	if strings.Index(number, "400") == 0 {
		return number
	} else {
		return "0" + number
	}
}

// 添加号码前置86，手机号前加 86，如：8613511112222
// 固话前加 86 加区号（首位 0 不写），如：
// 北京固话 861082325588
// 400 类电话前加 86，如：864003008000
func TelNoAdd86Head(number string) string {
	if len(number) == 0 {
		return number
	}

	if strings.Index(number, "86") == 0 {
		return number
	}

	if ruleUtil.CheckIsCnMobil(number) {
		return "86" + number
	}

	if strings.Index(number, "0") != 0 {
		return "86" + number
	} else {
		return "86" + number[1:]
	}
}

// 把aaa[bbb]解析成aaa bbb
func SplitMapFormatString(str string) (string, string, error) {
	str1 := strings.ReplaceAll(str, "]", "")
	str1 = strings.TrimSpace(str1)
	sarr := strings.Split(str1, "[")
	if len(sarr) != 2 {
		return "", "", fmt.Errorf("wrong format :string must just single []")
	}
	return sarr[0], sarr[1], nil
}

// NoCopyBytes2String b2s converts byte slice to a string without memory allocation.
// See https://groups.google.com/forum/#!msg/Golang-Nuts/ENgbUzYvCuU/90yGx7GUAgAJ .
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func NoCopyBytes2String(b []byte) string {
	if len(b) > 0 {
		/* #nosec G103 */
		return *(*string)(unsafe.Pointer(&b))
	} else {
		return ""
	}
}

// NoCopyString2Bytes s2b converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func NoCopyString2Bytes(s string) []byte {
	if len(s) > 0 {
		return unsafe.Slice(unsafe.StringData(s), len(s))
	} else {
		return nil
	}
}

//func NoCopyString2Bytes(s string) (b []byte) {
//	/* #nosec G103 */
//	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//	/* #nosec G103 */
//	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
//	bh.Data = sh.Data
//	bh.Len = sh.Len
//	bh.Cap = sh.Len
//	return b
//}

func ExtractNumberStrings(str string) []string {
	reg := regexp.MustCompile("[0-9]*\\.?[0-9]+")
	nodes := reg.FindAllString(str, -1)
	return nodes
}
