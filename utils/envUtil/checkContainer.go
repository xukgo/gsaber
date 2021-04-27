package envUtil

import (
	"fmt"
	"os"
	"strings"
)

func CheckIsContainerEnv() (bool, error) {
	//pid=1的应用就是容器
	pid := os.Getpid()
	if pid == 1 {
		return true, nil
	}

	filePath := "/proc/1/cgroup"
	contents, err := os.ReadFile(filePath)
	if err != nil {
		return false, fmt.Errorf("环境未能定位cgroup")
	}
	if len(contents) == 0 {
		return false, nil
	}

	arr := strings.Split(string(contents), "\n")
	dict := make(map[string]string)
	for _, s := range arr {
		arr1 := strings.Split(s, ":")
		if len(arr1) == 3 {
			key := arr1[1]
			if strings.Index(key, ",") < 0 {
				dict[key] = arr1[2]
			}
			arr2 := strings.Split(key, ",")
			for idx := range arr2 {
				dict[arr2[idx]] = arr1[2]
			}
		}
	}

	str, find := dict["devices"]
	if !find {
		return false, fmt.Errorf("环境未能定位cgroup.devices")
	}
	if checkContainerKey(str) {
		return true, nil
	}

	str, find = dict["cpu"]
	if !find {
		return false, fmt.Errorf("环境未能定位cgroup.cpu")
	}
	if checkContainerKey(str) {
		return true, nil
	}
	return false, nil
}

func checkContainerKey(str string) bool {
	if strings.Index(str, "docker") >= 0 || strings.Index(str, "Docker") >= 0 {
		return true
	}
	if strings.Index(str, "rkt") >= 0 || strings.Index(str, "Rkt") >= 0 {
		return true
	}
	if strings.Index(str, "sandbox") >= 0 || strings.Index(str, "Sandbox") >= 0 {
		return true
	}
	return false
}
