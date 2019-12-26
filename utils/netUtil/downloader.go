package netUtil

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"time"
)

func DownloadWebFile(durl string, savePath string, timeout time.Duration) error {
	uri, err := url.ParseRequestURI(durl)
	if err != nil {
		return err
	}

	filename := path.Base(uri.Path)
	fmt.Println("[*] Filename " + filename)

	client := http.DefaultClient
	if timeout > 0{
		client.Timeout = timeout //设置超时时间
	}

	resp, err := client.Get(durl)
	if err != nil {
		return err
	}

	//提醒一下这种的不支持断点续传
	if resp.ContentLength <= 0 {
		fmt.Println("[*] Destination server does not support breakpoint download.")
	}

	raw := resp.Body
	defer raw.Close()

	reader := bufio.NewReaderSize(raw, 1024 * 32)

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)

	var ret = true
	buff := make([]byte, 32*1024)
	written := 0
	go func() {
		for {
			nr, er := reader.Read(buff)
			if nr > 0 {
				nw, ew := writer.Write(buff[0:nr])
				if nw > 0 {
					written += nw
				}
				if ew != nil {
					err = ew
					break
				}
				if nr != nw {
					err = io.ErrShortWrite
					break
				}
			}
			if er != nil {
				if er != io.EOF {
					err = er
				}
				break
			}
		}
		if err != nil {
			ret = false
		}
	}()

	spaceTime := time.Second * 1
	ticker := time.NewTicker(spaceTime)
	lastWtn := 0
	stop := false

	for {
		select {
		case <-ticker.C:
			speed := written - lastWtn
			fmt.Printf("[*] Speed %s / %s \n", bytesToSize(speed), spaceTime.String())
			if written-lastWtn == 0 {
				ticker.Stop()
				stop = true
				break
			}
			lastWtn = written
		}
		if stop {
			if !ret{
				return err
			}else{
				return nil
			}
		}
	}

	return nil
}

func bytesToSize(length int) string {
	var k = 1024 // or 1024
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB"}
	if length == 0 {
		return "0 Bytes"
	}
	i := math.Floor(math.Log(float64(length)) / math.Log(float64(k)))
	r := float64(length) / math.Pow(float64(k), i)
	return strconv.FormatFloat(r, 'f', 3, 64) + " " + sizes[int(i)]
}