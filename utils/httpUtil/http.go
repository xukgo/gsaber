package httpUtil

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

const METHOD_GET = "GET"
const METHOD_POST = "POST"

var postJsonHeader = initPostJsonHeader()

func initPostJsonHeader() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json; charset=utf-8"
	return header
}

type Client struct {
	client *http.Client
	url    string
}

func NewClient(client *http.Client, url string) *Client {
	model := new(Client)
	model.client = client
	model.url = url
	return model
}

func (this *Client) Get() HttpResponse {
	resp, err := this.client.Get(this.url)
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	body, err := io.ReadAll(resp.Body)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	return NewHttpResponse(err, resp.StatusCode, resp.Status, body, resp.Header)
}

func (this *Client) PostJson(gson []byte) HttpResponse {
	return this.PostWithHeader(postJsonHeader, gson)
}

func (this *Client) PostWithHeader(header map[string]string, content []byte) HttpResponse {
	var reqs *http.Request
	var err error
	reqs, err = http.NewRequest(METHOD_POST, this.url, bytes.NewReader(content))

	if err != nil {
		return NewErrorHttpResponse(err)
	}

	if header != nil {
		for key, val := range header {
			reqs.Header[key] = []string{val}
		}
	}

	resp, err := this.client.Do(reqs)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	resBuff, err := io.ReadAll(resp.Body)
	return NewHttpResponse(err, resp.StatusCode, resp.Status, resBuff, resp.Header)
}

func (this *Client) HttpPostJsonWithFull(header map[string]string, cookies map[string]string, gson []byte) HttpResponse {
	var reqs *http.Request
	var err error
	if len(gson) == 0 {
		reqs, err = http.NewRequest(METHOD_POST, this.url, nil)
	} else {
		reqs, err = http.NewRequest(METHOD_POST, this.url, bytes.NewReader(gson))
	}

	if err != nil {
		return NewErrorHttpResponse(err)
	}

	if cookies != nil {
		cookieExpire := time.Now().Add(time.Hour * 12)
		for key, val := range cookies {
			//可以添加多个cookie
			cookie1 := &http.Cookie{Name: key, Value: val, HttpOnly: true, Expires: cookieExpire}
			reqs.AddCookie(cookie1)
		}
	}
	if header != nil {
		for key, val := range header {
			reqs.Header[key] = []string{val}
		}
	}

	resp, err := this.client.Do(reqs)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return NewErrorHttpResponse(err)
	}

	resBuff, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	return NewHttpResponse(err, resp.StatusCode, resp.Status, resBuff, resp.Header)
}
