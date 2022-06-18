package httpUploader

import (
	"bufio"
	"fmt"
	"github.com/valyala/fasthttp"
	"os"
	"testing"
	"time"
)

var singletonClient = &fasthttp.Client{
	// 读超时时间,不设置read超时,可能会造成连接复用失效
	ReadTimeout: time.Second * 30,
	// 写超时时间
	//WriteTimeout: time.Second * 30,
	// 60秒后，关闭空闲的活动连接
	MaxIdleConnDuration: time.Second * 60,
	// 当true时,从请求中去掉User-Agent标头
	NoDefaultUserAgentHeader: true,
	// 当true时，header中的key按照原样传输，默认会根据标准化转化
	DisableHeaderNamesNormalizing: true,
	//当true时,路径按原样传输，默认会根据标准化转化
	DisablePathNormalizing: true,
	Dial: (&fasthttp.TCPDialer{
		// 最大并发数，0表示无限制
		Concurrency: 10,
		// 将 DNS 缓存时间从默认分钟增加到一小时
		DNSCacheDuration: time.Second * 60,
	}).Dial,
}

func TestUploadNoLimit(t *testing.T) {
	//file, err := os.Open("F:\\FFOutput\\sample1.wav")
	file, err := os.Open("/home/hermes/Music/sample1.wav")
	if err != nil {
		t.FailNow()
		return
	}
	//不要忘记关闭打开的文件
	defer file.Close()
	reader := bufio.NewReader(file)
	cache := make([]byte, 100*1024)

	uploader := InitUploader(singletonClient, "http://192.168.5.164:8741", reader)
	uploader.SetCache(cache)
	uploader.SetFileName("sample.txt")
	uploader.AddFormValue("app_id", "appid123")
	uploader.AddFormValue("is_check_wav", "0")
	uploader.AddFormValue("createTime", time.Now().String())
	uploader.AddFormValue("file_type", "1")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = uploader.Upload(resp, time.Second*100)
	if err != nil {
		fmt.Printf("http post error:%s\n", err)
		t.FailNow()
	}
	fasthttp.ReleaseResponse(resp)
	fmt.Printf("http post success:%s\n", resp.Body())
}

func TestUploadLimitSpeed(t *testing.T) {
	file, err := os.Open("/home/hermes/Music/sample1.wav")
	//file, err := os.Open("F:\\FFOutput\\floop.wav")
	if err != nil {
		t.FailNow()
		return
	}
	//不要忘记关闭打开的文件
	defer file.Close()
	reader := bufio.NewReader(file)
	cache := make([]byte, 16*1024)

	uploader := InitUploader(singletonClient, "http://192.168.5.164:8741", reader)
	//uploader := InitUploader(singletonClient, "http://127.0.0.1:8741", reader)
	uploader.SetCache(cache)
	uploader.SetFileName("sample.txt")
	uploader.AddFormValue("app_id", "appid123")
	uploader.AddFormValue("is_check_wav", "0")
	uploader.AddFormValue("createTime", time.Now().String())
	uploader.AddFormValue("file_type", "1")
	uploader.SetRateBytes(10 * 1024 * 1024)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	startAt := time.Now()
	err = uploader.Upload(resp, time.Second*100)
	if err != nil {
		fmt.Printf("http post error:%s\n", err)
		time.Sleep(time.Hour)
		t.FailNow()
	}
	fmt.Printf("http post success:%s\n", resp.Body())
	totalWriteBytes := uploader.GetTotalWriteBytes()
	elapse := time.Since(startAt)
	speed := float64(totalWriteBytes) / 1024 / elapse.Seconds()
	fmt.Printf("send total=%d duration=%dms speed=%.2fkB/s\n", totalWriteBytes/1024, elapse.Milliseconds(), speed)
}
