package httpUploader

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"testing"
)

func TestGzipPipe(t *testing.T) {
	outputFile, _ := os.Create("/home/hermes/Music/sample1.gz")
	defer outputFile.Close()
	gzipWriter := gzip.NewWriter(outputFile) //gzipWriter:需要操作的句柄
	defer gzipWriter.Close()
	compress("/home/hermes/Music/sample1.wav", gzipWriter)
}

func compress(file string, gzipWriter *gzip.Writer) {
	readFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(readFile)
	cacheSize := 16 * 1024
	readBuff := make([]byte, cacheSize)
	writeBuff := make([]byte, cacheSize)

	piper, pipew := io.Pipe()
	defer piper.Close()

	fmt.Printf("begin compress\n")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			n, err := piper.Read(writeBuff)
			if n > 0 {
				n, err = gzipWriter.Write(writeBuff[:n])
				if err != nil {
					fmt.Printf("gzipWriter Write error:%s\n", err)
					return
				}
			}
			if err == io.EOF {
				fmt.Printf("end compress\n")
				return
			}
			if err != nil {
				gzipWriter.Flush()
				fmt.Printf("end compress\n")
				fmt.Printf("piper Read error:%s\n", err)
				return
			}
		}
	}()

	defer wg.Wait()
	for {
		s, e := reader.Read(readBuff)
		if s > 0 {
			pipew.Write(readBuff[:s])
		}
		if e == io.EOF {
			pipew.Close()
			break
		}
		if e != nil {
			fmt.Printf("file Read error:%s\n", err)
			pipew.Close()
			return
		}
	}
}
