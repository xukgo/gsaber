package httpPool

import (
	"crypto/tls"
	"net/http"
	"sync"
	"time"
)

type Pool struct {
	dict   map[string]*http.Client
	locker *sync.RWMutex
}

func NewPool() *Pool {
	repo := new(Pool)
	repo.locker = new(sync.RWMutex)
	repo.dict = make(map[string]*http.Client)
	return repo
}

func (this *Pool) DeleteClient(conf *HttpClientConfig) {
	var connUrl = conf.FormatConnectUrl()
	this.locker.Lock()
	defer this.locker.Unlock()
	if _, find := this.dict[connUrl]; find {
		return
	}
	delete(this.dict, connUrl)
}

func (this *Pool) GetClient(conf *HttpClientConfig) *http.Client {
	var connUrl = conf.FormatConnectUrl()
	this.locker.RLock()
	if pv, find := this.dict[connUrl]; find {
		this.locker.RUnlock()
		return pv
	}
	this.locker.RUnlock()

	this.locker.Lock()
	defer this.locker.Unlock()

	client := this.createClient(conf)
	this.dict[connUrl] = client
	return client
}
func (this *Pool) createClient(conf *HttpClientConfig) *http.Client {
	client := &http.Client{
		Timeout: conf.Timeout, //设置超时时间
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},

			//Dial: (&net.Dialer{
			//	Timeout:   3 * time.Second, //限制建立TCP连接的时间
			//	KeepAlive: time.Duration(conf.maxAliveCount) * time.Second,
			//}).Dial,
			//Dial:                  dialer.Dial,
			TLSHandshakeTimeout: 3 * time.Second, //限制 TLS握手的时间
			//ResponseHeaderTimeout: 3 * time.Second,                                 //限制读取response header的时间,默认 timeout + 5*time.Second
			ExpectContinueTimeout: 3 * time.Second,          //限制client在发送包含 Expect: 100-continue的header到收到继续发送body的response之间的时间等待。
			MaxIdleConns:          conf.MaxIdleConns,        //所有host的连接池最大连接数量，默认无穷大
			MaxIdleConnsPerHost:   conf.MaxIdleConnsPerHost, //每个host的连接池最大空闲连接数,默认2
			MaxConnsPerHost:       conf.MaxConnsPerHost,     //每个host的最大连接数量
			IdleConnTimeout:       conf.IdleConnTimeout,     //how long an idle connection is kept in the connection pool.

		},
	}
	return client
}
