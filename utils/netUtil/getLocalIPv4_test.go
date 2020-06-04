package netUtil

import (
	"fmt"
	"testing"
)

func Test_getLocalIPv4(t *testing.T) {
	ipArr, err := GetIPv4("private", []string{"172."})
	if err != nil {
		t.FailNow()
	}
	fmt.Println(ipArr)
}
