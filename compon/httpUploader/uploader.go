package httpUploader

import (
	"github.com/juju/ratelimit"
	"github.com/valyala/fasthttp"
	"github.com/xukgo/gsaber/utils/limitio"
	"io"
	"mime/multipart"
	"path"
	"sync/atomic"
	"time"
)

type Uploader struct {
	formArgs          map[string]string
	fileFieldName     string
	fileName          string
	reader            io.Reader
	cache             []byte
	url               string
	client            *fasthttp.Client
	continue100Enable bool
	rateBytes         int           //每秒传输的字节数,0或者负数表示不限速
	limitPrecision    time.Duration //限速精度，精度越大越不准可能超速，精度越小越不准可能欠速，最多1秒，最少10ms，默认250ms
	totalWriteBytes   int64
}

func InitUploader(client *fasthttp.Client, url string, reader io.Reader) Uploader {
	model := Uploader{
		formArgs:          make(map[string]string),
		fileFieldName:     "file",
		reader:            reader,
		client:            client,
		url:               url,
		continue100Enable: false,
		rateBytes:         0,
		limitPrecision:    time.Millisecond * 250,
		totalWriteBytes:   0,
	}
	return model
}

func (this *Uploader) GetTotalWriteBytes() int64 {
	n := atomic.LoadInt64(&this.totalWriteBytes)
	return n
}

func (this *Uploader) SetContinue100Enable(enable bool) {
	this.continue100Enable = enable
}
func (this *Uploader) SetCache(cache []byte) {
	this.cache = cache
}
func (this *Uploader) SetRateBytes(rate int) {
	this.rateBytes = rate
}
func (this *Uploader) SetLimitPrecision(lp time.Duration) {
	if lp > time.Second {
		lp = time.Second
	} else if lp < time.Millisecond*10 {
		lp = time.Millisecond * 10
	}
	this.limitPrecision = lp
}

// SetFileName 设置 filename="xxx.xx"
func (this *Uploader) SetFileName(name string) {
	this.fileName = name
}
func (this *Uploader) AddFormValue(key, value string) {
	this.formArgs[key] = value
}

// SetFieldFileName 这个一般不用设置，默认file
func (this *Uploader) SetFieldFileName(name string) {
	this.fileFieldName = name
}

func (this *Uploader) Upload(response *fasthttp.Response, timeout time.Duration) error {
	piper, pipew := io.Pipe()
	defer piper.Close()

	var buff = this.cache
	if len(buff) == 0 {
		buff = make([]byte, 10*1024)
	}

	var limitWriter *limitio.LimitWriter = nil
	if this.rateBytes > 0 {
		segLimit := int64(float64(this.rateBytes) * float64(this.limitPrecision) / float64(time.Second))
		if segLimit < 1 {
			segLimit = 1
		}
		bucket := ratelimit.NewBucketWithQuantum(this.limitPrecision, segLimit, segLimit)
		limitWriter = limitio.NewLimitWriter(pipew, bucket)
	} else {
		limitWriter = limitio.NewLimitWriter(pipew, nil)
	}
	multipartWriter := multipart.NewWriter(limitWriter)
	contentType := multipartWriter.FormDataContentType()

	//wg := &sync.WaitGroup{}
	//wg.Add(1)

	go func() {
		//defer wg.Done()
		defer pipew.Close()

		this.limitWrite(multipartWriter, buff)
	}()

	//构建request，发送请求
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	request.Header.SetContentType(contentType)
	request.Header.SetMethod("POST")
	if this.continue100Enable {
		request.Header.Set("Expect", "100-continue")
	}
	request.SetBodyStream(piper, -1)
	request.SetRequestURI(this.url)

	err := this.client.DoTimeout(request, response, timeout)
	if err != nil {
		return err
	}
	//wg.Wait()
	if limitWriter != nil {
		atomic.StoreInt64(&this.totalWriteBytes, limitWriter.GetTotalWriteCount())
	}
	return nil
}

func (this *Uploader) limitWrite(multipartWriter *multipart.Writer, buff []byte) {
	//创建一个multipart文件写入器，方便按照http规定格式写入内容
	for k, v := range this.formArgs {
		err := multipartWriter.WriteField(k, v)
		if err != nil {
			return
		}
	}
	fileName := this.fileName
	if len(fileName) == 0 {
		fileName = "sample.txt"
	} else {
		fileName = path.Base(this.fileName)
	}
	//从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := multipartWriter.CreateFormFile(this.fileFieldName, path.Base(this.fileName))
	if err != nil {
		return
	}
	defer multipartWriter.Close()

	for {
		n, err := this.reader.Read(buff)
		if n > 0 {
			_, err = fileWriter.Write(buff[:n])
			if err != nil {
				//fmt.Printf("fileWriter Write error:%s\n", err)
				return
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			//fmt.Printf("bufio read file error:%s\n", err)
			return
		}
	}
}
