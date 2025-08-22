/*
@Time : 2019/10/19 10:47
@Author : Hermes
@File : rule_test
@Description:
*/
package ruleUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilEqual(t *testing.T) {
	var br bool
	var a map[string]int = nil
	br = InterfaceIsNil(a)
	if !br {
		t.FailNow()
	}
	var c []string = nil
	br = InterfaceIsNil(c)
	if !br {
		t.FailNow()
	}
}
func TestFloatSuccess(t *testing.T) {
	if !CheckIsFloat("0") {
		t.Fail()
	}
	if !CheckIsFloat("1") {
		t.Fail()
	}
	if !CheckIsFloat("1234567890") {
		t.Fail()
	}
	if !CheckIsFloat("+0") {
		t.Fail()
	}
	if !CheckIsFloat("+1") {
		t.Fail()
	}
	if !CheckIsFloat("+1234567890") {
		t.Fail()
	}
	if !CheckIsFloat("-1") {
		t.Fail()
	}
	if !CheckIsFloat("-46435435465390") {
		t.Fail()
	}

	if !CheckIsFloat("0.135646384") {
		t.Fail()
	}
	if !CheckIsFloat("1.48346384") {
		t.Fail()
	}
	if !CheckIsFloat("1234567890.46876") {
		t.Fail()
	}
	if !CheckIsFloat("+0.564684") {
		t.Fail()
	}
	if !CheckIsFloat("+1.1463846834") {
		t.Fail()
	}
	if !CheckIsFloat("+1234567890.468494") {
		t.Fail()
	}
	if !CheckIsFloat("-1.46846874") {
		t.Fail()
	}
	if !CheckIsFloat("-46435435465390.54864684") {
		t.Fail()
	}
}

func TestFloatFail(t *testing.T) {
	if CheckIsFloat("0a") {
		t.Fail()
	}
	if CheckIsFloat("144f4fd4sf4") {
		t.Fail()
	}
	if CheckIsFloat("123d456g7890") {
		t.Fail()
	}
	if CheckIsFloat("+0+") {
		t.Fail()
	}
	if CheckIsFloat("+1)") {
		t.Fail()
	}
	if CheckIsFloat("+&1234f567890") {
		t.Fail()
	}
	if CheckIsFloat("#-1") {
		t.Fail()
	}
	if CheckIsFloat("!-46435435465390") {
		t.Fail()
	}

	if CheckIsFloat("*0.135646384") {
		t.Fail()
	}
	if CheckIsFloat("1..48346384") {
		t.Fail()
	}
	if CheckIsFloat("1.234567890.46876") {
		t.Fail()
	}
	if CheckIsFloat("+0.56+4684") {
		t.Fail()
	}
	if CheckIsFloat("+1.14638-46834") {
		t.Fail()
	}
	if CheckIsFloat("+123456/7890.468494") {
		t.Fail()
	}
	if CheckIsFloat("-1.4684$6874") {
		t.Fail()
	}
	if CheckIsFloat("-46435435~465390.54864684") {
		t.Fail()
	}
}

func Benchmark_intRule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckIsIntRange("100", 0, 65535)
	}
}

func TestCheckIsLenPhoneNumber(t *testing.T) {
	var no string
	no = "+8615986400521"
	if !CheckIsLenPhoneNumber(no, 20) {
		t.Fail()
	}
	no = "-8615986400521"
	if CheckIsLenPhoneNumber(no, 20) {
		t.Fail()
	}
	no = "8615986400521"
	if !CheckIsLenPhoneNumber(no, 20) {
		t.Fail()
	}
	no = "15986400521"
	if !CheckIsLenPhoneNumber(no, 20) {
		t.Fail()
	}
}

func TestCheckIsCnPhoneNumber(t *testing.T) {
	var no string
	no = "+8615986400521"
	if !CheckIsCnMobilWith86Start(no) {
		t.Fail()
	}
	no = "8615986400521"
	if !CheckIsCnMobilWith86Start(no) {
		t.Fail()
	}
	no = "-8615986400521"
	if CheckIsCnMobilWith86Start(no) {
		t.Fail()
	}
}

func Test_CheckIpPort(t *testing.T) {
	tcs := []string{
		"192.168.1.1:8080",
		"[2001:db8::1]:8080",
		"127.0.0.1:80",
		"invalid:port",
		"256.256.256.256:80",
	}

	assert.True(t, IsValidIPPort(tcs[0]))
	assert.True(t, IsValidIPPort(tcs[1]))
	assert.True(t, IsValidIPPort(tcs[2]))
	assert.True(t, !IsValidIPPort(tcs[3]))
	assert.True(t, !IsValidIPPort(tcs[4]))
}
