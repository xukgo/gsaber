package fileUtil

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//判断文件是否存在
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	br, _ := Exists(path)
	if !br {
		return false
	}
	return !IsDir(path)
}

func SaveFile(path, content string) error {
	if strings.Contains(path, "/") {

		paths := strings.Split(path, "/")
		dir := strings.Join(paths[0:len(paths)-1], "/")
		if !IsDir(dir) {
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return err
			}
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.WriteString(content)
	defer f.Close()

	return nil
}

func ReadFile(path string) (string, error) {
	if IsFile(path) {
		d, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(d), nil
	}
	return "", fmt.Errorf("%s is not exists", path)
}

//获取文件的后缀，不包含点
func GetFileExt(fileName string) string {
	index := strings.LastIndex(fileName, ".")
	if index <= 0 || index == len(fileName)-1 {
		return ""
	}

	return fileName[index+1:]
}

//获取文件的文件夹路径，包含/
func GetFileDirUrl(fileName string) string {
	index := strings.LastIndex(fileName, "/")
	if index <= 0 {
		return ""
	}

	return fileName[:index+1]
}

//获取文件路径的文件名
func GetFileName(path string) string {
	index := strings.LastIndex(path, "/")
	if index < 0 {
		return path
	}

	return path[index+1:]
}

//获取文件路径的文件名，不包含后缀
func GetFileNameNoExt(path string) string {
	var fileName string
	index := strings.LastIndex(path, "/")
	if index < 0 {
		fileName = path
	} else {
		fileName = path[index+1:]
	}

	index = strings.LastIndex(fileName, ".")
	if index < 0 {
		return fileName
	} else {
		return fileName[:index]
	}
}
