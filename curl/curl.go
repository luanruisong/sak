package utils

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Curl struct {
	Domain string
}

var (
	transport = http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          200,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   50,
	}

	httpClient = http.Client{
		Transport: &transport,
		Timeout:   5 * time.Second,
	}
)

// Get http-get请求
func (c Curl) Get(uri string, params map[string]string) ([]byte, error) {
	return c.GetWithTimout(uri, params, 0)
}

func (c Curl) GetWithTimout(uri string, params map[string]string, timeout time.Duration) ([]byte, error) {
	parseURL, parseErr := url.Parse(c.Domain + uri)
	if parseErr != nil {
		return []byte{}, parseErr
	}
	if len(params) > 0 {
		q := parseURL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		parseURL.RawQuery = q.Encode()
	}
	requestURL := parseURL.String()
	var (
		client http.Client
	)
	if timeout > 0 {
		client = http.Client{
			Transport: &transport,
			Timeout:   timeout,
		}
	} else {
		client = httpClient
	}
	result, err := client.Get(requestURL)
	if err != nil {
		return []byte{}, err
	}
	resultByte, err1 := ioutil.ReadAll(result.Body)
	defer result.Body.Close()
	if err1 != nil {
		return []byte{}, err1
	}
	return resultByte, nil
}
