// +build linux

package fileUtil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
获取程序运行路径，不会穿透软连接
*/
func GetCurrentDirectory() string {
	//var dir string
	//ex, err := os.Executable()
	//if err != nil {
	//	fmt.Println(err)
	//	return ""
	//}
	//
	//dir = filepath.Dir(ex)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

/*
获取程序运行路径，穿透软连接
*/
func GetRealRunDirectory() string {
	var dir string
	ex, err := os.Executable()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	exReal, err := filepath.EvalSymlinks(ex)
	if err != nil {
		panic(err)
	}
	dir = filepath.Dir(exReal)
	//fmt.Println(dir)
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
