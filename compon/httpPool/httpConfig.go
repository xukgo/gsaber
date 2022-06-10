package httpPool

import (
	"fmt"
	"net/http"
	"time"
)

type OptionFunc func(*ClientOption)

type ClientOption struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	MaxConnsPerHost     int
	IdleConnTimeout     time.Duration
}

type HttpClientConfig struct {
	ClientOption
	Domain  string
	Timeout time.Duration
}

func InitHttpClientConfig(domain string, timeout time.Duration, options ...OptionFunc) HttpClientConfig {
	conf := HttpClientConfig{
		Domain:  domain,
		Timeout: timeout,
	}
	conf.MaxIdleConns = 0
	conf.MaxIdleConnsPerHost = http.DefaultMaxIdleConnsPerHost
	conf.MaxConnsPerHost = 0
	conf.IdleConnTimeout = time.Second * 10

	for _, op := range options {
		op(&conf.ClientOption)
	}
	return conf
}

func (this HttpClientConfig) FormatConnectUrl() string {
	return fmt.Sprintf("%s#%d#%d#%d#%d#%d", this.Domain, int64(this.Timeout), this.MaxIdleConns,
		this.MaxIdleConnsPerHost, this.MaxConnsPerHost, this.IdleConnTimeout)
}

func WithMaxIdleConns(MaxIdleConns int) OptionFunc {
	return func(option *ClientOption) {
		option.MaxIdleConns = MaxIdleConns
	}
}

func WithMaxIdleConnsPerHost(MaxIdleConnsPerHost int) OptionFunc {
	return func(option *ClientOption) {
		option.MaxIdleConnsPerHost = MaxIdleConnsPerHost
	}
}

func WithMaxConnsPerHost(MaxConnsPerHost int) OptionFunc {
	return func(option *ClientOption) {
		option.MaxConnsPerHost = MaxConnsPerHost
	}
}

func WithIdleConnTimeout(IdleConnTimeout time.Duration) OptionFunc {
	return func(option *ClientOption) {
		option.IdleConnTimeout = IdleConnTimeout
	}
}
