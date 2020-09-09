package netUtil

import (
	"fmt"
	"net"
	"strings"
)

const IPV4_LOCALHOST = "127.0.0.1"

func GetNetIpList() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ipv4List = make([]net.IP, 0, len(addrs))
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && ipnet.IP.To4() != nil {
			ipv4List = append(ipv4List, ipnet.IP.To4())
		}
	}
	return ipv4List, nil
}

//理论上的localhost地址
func CheckIsLocalhost(ip net.IP) bool {
	return ip.IsLoopback()
}

//理论上的私网地址
func CheckIsPrivateIP(ip net.IP) bool {
	if ip[12] == 10 {
		return true
	}
	if ip[12] == 172 && ip[13] >= 16 && ip[13] <= 31 {
		return true
	}
	if ip[12] == 192 && ip[13] == 168 {
		return true
	}
	return false
}

func checkIpPrefixMatch(ip net.IP, prefixArr []string) bool {
	if len(prefixArr) == 0 {
		return true
	}

	ipStr := ip.String()
	for i := range prefixArr {
		if strings.HasPrefix(ipStr, prefixArr[i]) {
			return true
		}
	}
	return false
}

func CheckIsPublicIP(ip net.IP) bool {
	match := ip.IsGlobalUnicast()
	if !match {
		return false
	}
	match = CheckIsPrivateIP(ip)
	if match {
		return false
	}
	return true
}

func GetIPv4(ipType string, filterPrefixList []string) ([]string, error) {
	var resIpArr = make([]string, 0, 1)
	if ipType == "local" || ipType == "localhost" || ipType == "loopback" {
		resIpArr = append(resIpArr, IPV4_LOCALHOST)
		return resIpArr, nil
	}

	ipv4List, err := GetNetIpList()
	if err != nil {
		return nil, err
	}

	if ipType == "private" {
		for i := range ipv4List {
			if CheckIsPrivateIP(ipv4List[i]) && checkIpPrefixMatch(ipv4List[i], filterPrefixList) {
				resIpArr = append(resIpArr, ipv4List[i].String())
			}
		}
		return resIpArr, nil
	}

	if ipType == "public" {
		for i := range ipv4List {
			if CheckIsPublicIP(ipv4List[i]) && checkIpPrefixMatch(ipv4List[i], filterPrefixList) {
				resIpArr = append(resIpArr, ipv4List[i].String())
			}
		}
		return resIpArr, nil
	}

	return nil, fmt.Errorf("unknow get type %s", ipType)
}
