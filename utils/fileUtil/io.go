package fileUtil

import (
	"bufio"
	"io"
	"os"
)

func ReadFileByLine(fileName string, handler func(string) bool) error {
	return ReadFileByDelim(fileName, '\n', handler)
}

func ReadFileByDelim(fileName string, endSeg byte, handler func(string) bool) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString(endSeg)
		if !handler(line) {
			return nil
		}

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
