/*
@Time : 2019/10/19 11:07
@Author : Hermes
@File : strings_test
@Description:
*/
package stringUtil

import (
	"fmt"
	"strings"
	"testing"
)

func TestStartWithSuccess(t *testing.T) {
	if !StartWith("abcdefg", "a") {
		t.Fail()
	}
	if !StartWith("abcdefg", "ab") {
		t.Fail()
	}
	if !StartWith("abcdefg", "abc") {
		t.Fail()
	}
	if !StartWith("abcdefg", "abcdefg") {
		t.Fail()
	}
}

func TestStartWithFailed(t *testing.T) {
	if StartWith("1234abcdefg", "124") {
		t.Fail()
	}
	if StartWith("1234abcdefg", "2") {
		t.Fail()
	}
	if StartWith("1234abcdefg", "abc") {
		t.Fail()
	}
	if StartWith("1234abcdefg", "ab") {
		t.Fail()
	}
	if StartWith("abcdefg", "bc") {
		t.Fail()
	}
}

func TestStartWithIndexSuccess(t *testing.T) {
	if !StartWithIndex("abcdefg", 0, "a") {
		t.Fail()
	}
	if !StartWithIndex("abcdefg", 1, "bc") {
		t.Fail()
	}
	if !StartWithIndex("abcdefg", 2, "cde") {
		t.Fail()
	}
	if !StartWithIndex("abcdefg", 4, "efg") {
		t.Fail()
	}
}

func TestStartWithIndexFailed(t *testing.T) {
	if StartWithIndex("abcdefg", 1, "a") {
		t.Fail()
	}
	if StartWithIndex("abcdefg", 2, "bc") {
		t.Fail()
	}
	if StartWithIndex("abcdefg", 3, "cde") {
		t.Fail()
	}
	if StartWithIndex("abcdefg", 0, "efg") {
		t.Fail()
	}
}

func TestSplitContainsSeps1(t *testing.T) {
	exp := "1+23+456-7890*(-45635/(454*468))"
	actual := SplitContainsSeps(exp, []string{"+", "-", "*", "/", "(", ")"})
	fmt.Println(actual)
	if len(actual) != 18 {
		t.Fail()
	}
	if actual[0] != "1" {
		t.Fail()
	}
	if actual[1] != "+" {
		t.Fail()
	}
	if actual[2] != "23" {
		t.Fail()
	}
	if actual[3] != "+" {
		t.Fail()
	}
	if actual[4] != "456" {
		t.Fail()
	}
	if actual[5] != "-" {
		t.Fail()
	}
	if actual[6] != "7890" {
		t.Fail()
	}
}

func TestSplitContainsSeps2(t *testing.T) {
	exp := "1+23+456-7890"
	actual := SplitContainsSeps(exp, []string{"+", "-", "*", "/", "(", ")"})
	fmt.Println(actual)
}

func Benchmark_StartWith1(b *testing.B){
	for i:=0;i<b.N;i++{
		StartWith("1234abcdefg", "1234abcdefa")
	}
}

func Benchmark_OfficeStartWith1(b *testing.B){
	for i:=0;i<b.N;i++{
		strings.HasPrefix("1234abcdefg", "1234abcdefa")
	}
}