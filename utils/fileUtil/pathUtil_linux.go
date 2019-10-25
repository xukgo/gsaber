// +build linux

package fileUtil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetAbsUrl(relaPath string) string {
	if strings.Index(relaPath, "/") == 0 {
		return relaPath
	} else {
		currentPath := GetCurrentDirectory()
		return filepath.Join(currentPath, relaPath)
	}
}
