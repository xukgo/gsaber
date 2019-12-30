/*
@Time : 2019/9/24 20:16
@Author : Hermes
@File : string
@Description:
*/
package ruleUtil

import (
	"regexp"
	"strconv"
	"time"
)

//判断是否是网络链接，检查是否是http://或者https://开头
func CheckIsWebUrl(str string) bool {
	matched, _ := regexp.MatchString(`^http(s)?://.+`, str)
	return matched
}

//判断是否是浮点数字符串,允许+-符号，支持没有小数点的浮点数
func CheckIsFloat(str string) bool {
	matched, _ := regexp.MatchString(`^[-+]?[0-9]*\.?[0-9]+$`, str)
	return matched
}

//判断是否是数字字符串,允许0开头
func CheckIsInteger(str string) bool {
	matched, _ := regexp.MatchString(`^-?[0-9]\d*$`, str)
	return matched
}

//判断是否是非负数字符串,允许0开头
func CheckIsNonNegativeInteger(str string) bool {
	matched, _ := regexp.MatchString(`^[0-9]\d*$`, str)
	return matched
}

//是否是在一个范围内的数字
func CheckIsIntRange(str string, min int, max int) bool {
	if !CheckIsInteger(str) {
		return false
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	if val < min {
		return false
	}
	if val >= max {
		return false
	}
	return true
}

//判断是否是合法的电话，只判断是纯数字，并且长度不超过16
func CheckIsPhoneNumber(str string) bool {
	if len(str) > 16 {
		return false
	}

	matched, _ := regexp.MatchString(`^\d+`, str)
	return matched
}

//判断是否是合法的电话，只判断是纯数字，并且长度不超过指定
func CheckIsLenPhoneNumber(str string, length int) bool {
	if len(str) > length {
		return false
	}

	matched, _ := regexp.MatchString(`^\d+`, str)
	return matched
}

//判断是否是手机号,简单点，1开头，然后一共11位数字
func CheckIsCnMobil(str string) bool {
	matched, _ := regexp.MatchString(`^1\d{10}$`, str)
	return matched
}

func CheckIsCnMobilWith86Start(str string) bool {
	matched, _ := regexp.MatchString(`^861\d{10}$`, str)
	return matched
}

//是否是时间戳，以秒计数
func CheckIsTimestampSecond(str string) bool {
	uc, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if uc < 1000000000 || uc > 5000000000 {
		return false
	}
	return true
}

//是否是时间戳，以毫秒计数
func CheckIsTimestampMillisecond(str string) bool {
	uc, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if uc < 1000000000000 || uc > 5000000000000 {
		return false
	}
	return true
}

//是否2019-01-01 20:00:00这样的时间表达式允许毫秒加在后面
func CheckIsBaseFormatTime(str string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", str)
	if err != nil {
		return false
	}
	return true
}

//是否20190101200000这样的时间表达式，这个后面没有毫秒
func CheckIsTightShortFormatTime(str string) bool {
	_, err := time.Parse("20060102150405", str)
	if err != nil {
		return false
	}
	return true
}
