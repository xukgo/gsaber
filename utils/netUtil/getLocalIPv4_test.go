package netUtil

import (
	"fmt"
	"net"
	"testing"
)

func Test_getLocalIPv4_private1(t *testing.T) {
	ipArr, err := GetIPv4("private", []string{"172."})
	if err != nil {
		t.FailNow()
	}
	fmt.Println(ipArr)
}

func Test_privateIP(t *testing.T) {
	ip := net.ParseIP("192.168.1.10")
	match := CheckIsPrivateIP(ip)
	if !match {
		t.FailNow()
	}
	ip = net.ParseIP("172.16.10.24")
	match = CheckIsPrivateIP(ip)
	if !match {
		t.FailNow()
	}
	ip = net.ParseIP("10.16.10.24")
	match = CheckIsPrivateIP(ip)
	if !match {
		t.FailNow()
	}

	ip = net.ParseIP("120.77.204.33")
	match = CheckIsPrivateIP(ip)
	if match {
		t.FailNow()
	}
}

func Test_publicIP(t *testing.T) {
	ip := net.ParseIP("172.16.10.24")
	match := CheckIsPublicIP(ip)
	if match {
		t.FailNow()
	}
	ip = net.ParseIP("120.77.204.33")
	match = CheckIsPublicIP(ip)
	if !match {
		t.FailNow()
	}
}
