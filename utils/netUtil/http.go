package netUtil

import (
	"bytes"
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

const METHOD_GET = "GET"
const METHOD_POST = "POST"
const METHOD_PUT = "PUT"
const METHOD_DELETE = "DELETE"

type HttpResponse struct {
	Error      error
	Data       []byte
	StatusCode int    // e.g. 200
	Status     string // e.g. "200 OK"
}

func NewHttpResponse(err error, code int, state string, data []byte) HttpResponse {
	model := HttpResponse{
		Error:      err,
		Data:       data,
		StatusCode: code,
		Status:     state,
	}
	return model
}

func NewErrorHttpResponse(err error) HttpResponse {
	model := HttpResponse{
		Error: err,
	}
	return model
}

func HttpGet(url string, timeout time.Duration) HttpResponse {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	resp, err := client.Get(url)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	return NewHttpResponse(err, resp.StatusCode, resp.Status, body)
}

func HttpPostJson(url string, gson string, timeout int, multiplex bool) HttpResponse {
	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=utf-8"
	return HttpPostJsonWithHeader(url, header, gson, timeout, multiplex)
}

func HttpPostJsonWithHeader(url string, header map[string]string, gson string, timeout int, multiplex bool) HttpResponse {
	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}

	//如果是https，跳过ssl验证
	if strings.Index(url, "https") >= 0 {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	var reqs *http.Request
	var err error
	if gson == "" {
		reqs, err = http.NewRequest(METHOD_POST, url, nil)
	} else {
		reqs, err = http.NewRequest(METHOD_POST, url, strings.NewReader(gson))
	}

	if err != nil {
		return NewErrorHttpResponse(err)
	}

	if multiplex {
		reqs.Close = false
	} else {
		reqs.Close = true
	}

	if header != nil {
		for key, val := range header {
			reqs.Header.Add(key, val)
		}
	}

	resp, err := client.Do(reqs)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	resBuff, err := io.ReadAll(resp.Body)
	return NewHttpResponse(err, resp.StatusCode, resp.Status, resBuff)
}

func HttpPostFileAndDataWithHeader(url string, header map[string]string, fieldDict map[string]string,
	fileFieldName string, filePath string, timeout int) HttpResponse {
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)

	//fmt.Printf("HttpPostFileAndDataWithAuth:url[%v] auth[%v]\r\n", url, authString)
	fileinfo, err := os.Stat(filePath)
	if err != nil {
		return NewErrorHttpResponse(err)
	}
	fileName := fileinfo.Name()

	fileWriter, _ := bodyWriter.CreateFormFile(fileFieldName, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	io.Copy(fileWriter, file)

	if fieldDict != nil {
		for key, val := range fieldDict {
			_ = bodyWriter.WriteField(key, val)
		}
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	reqs, err := http.NewRequest("POST", url, bodyBuffer)
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	reqs.Header.Add("Content-Type", contentType)
	if header != nil {
		for key, val := range header {
			reqs.Header.Add(key, val)
		}
	}

	client := &http.Client{Timeout: time.Duration(timeout) * time.Millisecond}
	resp, err := client.Do(reqs)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	resBuff, err := io.ReadAll(resp.Body)
	return NewHttpResponse(err, resp.StatusCode, resp.Status, resBuff)
}
