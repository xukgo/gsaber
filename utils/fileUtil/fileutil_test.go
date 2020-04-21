package fileUtil

import "testing"

func TestGetFileName(t *testing.T) {
	fileName := GetFileName("/usr/local/aaa.wav")
	if fileName != "aaa.wav" {
		t.Fail()
	}
	fileName = GetFileName("bbb.wav")
	if fileName != "bbb.wav" {
		t.Fail()
	}
	fileName = GetFileName("/usr/local/")
	if fileName != "" {
		t.Fail()
	}
}
