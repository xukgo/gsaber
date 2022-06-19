package httpUploader

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"os"
	"sync"
	"sync/atomic"
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
	cache := make([]byte, 16*1024)

	uploader := InitUploader(singletonClient, "http://192.168.5.164:8741", func(fileWriter io.Writer) error {
		n, err := reader.Read(cache)
		if n > 0 {
			_, err = fileWriter.Write(cache[:n])
			if err != nil {
				return err
			}
		}
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		return nil
	})
	//uploader.SetCache(cache)
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
	cache := make([]byte, 8*1024*1024)

	var gzReader = NewGzReadWriter(reader)
	defer gzReader.Close()
	once := new(sync.Once)

	uploader := InitUploader(singletonClient, "http://192.168.5.164:8741", func(fileWriter io.Writer) error {
		once.Do(func() {
			gzReader.SetWriter(fileWriter)
		})
		n, err := gzReader.ReadWrite(cache)
		if n > 0 {
			_, err = fileWriter.Write(cache[:n])
			if err != nil {
				return err
			}
		}
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}
		return nil
	})
	//uploader := InitUploader(singletonClient, "http://127.0.0.1:8741", reader)
	//uploader.SetCache(cache)
	uploader.SetFileName("sample.txt")
	uploader.AddFormValue("app_id", "appid123")
	uploader.AddFormValue("is_check_wav", "0")
	uploader.AddFormValue("createTime", time.Now().String())
	uploader.AddFormValue("file_type", "1")
	//uploader.SetRateBytes(50 * 1024 * 1024)

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

	compressFactor := float64(gzReader.compressSize) / float64(gzReader.sourceSize)
	fmt.Printf("compress rate:%.2f src:%d target:%d\n", compressFactor, gzReader.sourceSize, gzReader.compressSize)
}

type GzReadWriter struct {
	reader       io.Reader
	gzWriter     *gzip.Writer
	sourceSize   int64
	compressSize int64
}

func NewGzReadWriter(reader io.Reader) *GzReadWriter {
	return &GzReadWriter{reader: reader}
}

func (this *GzReadWriter) SetWriter(writer io.Writer) error {
	gzWriter, err := gzip.NewWriterLevel(writer, gzip.DefaultCompression)
	if err != nil {
		return err
	}
	this.gzWriter = gzWriter
	return nil
}

func (this *GzReadWriter) GetSourceSize() int64 {
	return atomic.LoadInt64(&this.sourceSize)
}

func (this *GzReadWriter) GetCompressSize() int64 {
	return atomic.LoadInt64(&this.compressSize)
}
func (this *GzReadWriter) Close() error {
	if this.gzWriter != nil {
		return this.gzWriter.Close()
	}
	return nil
}

func (this *GzReadWriter) ReadWrite(p []byte) (n int, err error) {
	if this.gzWriter == nil {
		return 0, fmt.Errorf("gzWriter is nil")
	}
	n, err = this.reader.Read(p)
	if err == io.EOF && n > 0 {
		atomic.AddInt64(&this.sourceSize, int64(n))
		n, err = this.gzWriter.Write(p[:n])
		atomic.AddInt64(&this.compressSize, int64(n))
		return n, io.EOF
	}
	if err != nil {
		return 0, err
	}
	if n > 0 {
		atomic.AddInt64(&this.sourceSize, int64(n))
		n, err = this.gzWriter.Write(p[:n])
		atomic.AddInt64(&this.compressSize, int64(n))
		return n, err
	}
	return 0, nil
}
