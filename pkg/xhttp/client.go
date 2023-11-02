package xhttp

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
	bodySize   int // body size limit(MB), default is 10MB
}

func defaultClient() *Client {
	return &Client{
		HttpClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: defaultTransportDialContext(&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
					DualStack: true,
				}),
				TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				DisableKeepAlives:     true,
				ForceAttemptHTTP2:     true,
			},
		},
		bodySize: 10, // default is 10MB
	}
}

// NewClient , default tls.Config{InsecureSkipVerify: true}
func NewClient() (client *Client) {
	return defaultClient()
}

func (c *Client) SetTransport(transport *http.Transport) (client *Client) {
	c.HttpClient.Transport = transport
	return c
}

func (c *Client) SetTLSConfig(tlsCfg *tls.Config) (client *Client) {
	c.HttpClient.Transport.(*http.Transport).TLSClientConfig = tlsCfg
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) (client *Client) {
	c.HttpClient.Timeout = timeout
	return c
}

// set body size (MB), default is 10MB
func (c *Client) SetBodySize(sizeMB int) (client *Client) {
	c.bodySize = sizeMB
	return c
}

// typeStr is request type and response type
// default is TypeJSON
// first param is request type
// second param is response data type
func (c *Client) Req(typeStr ...string) *Request {
	var (
		reqTp = TypeJSON    // default
		resTp = ResTypeJSON // default
		tLen  = len(typeStr)
	)
	switch {
	case tLen == 1:
		tpp := typeStr[0]
		if _, ok := _ReqContentTypeMap[tpp]; ok {
			reqTp = tpp
		}
	case tLen > 1:
		// first param is request type
		tpp := typeStr[0]
		if _, ok := _ReqContentTypeMap[tpp]; ok {
			reqTp = tpp
		}
		// second param is response data type
		stpp := typeStr[1]
		if _, ok := _ResTypeMap[stpp]; ok {
			resTp = stpp
		}
	}
	if c == nil {
		c = defaultClient()
	}
	r := &Request{
		client:       c,
		Header:       make(http.Header),
		requestType:  reqTp,
		responseType: resTp,
	}
	r.Header.Set("Content-Type", _ReqContentTypeMap[reqTp])
	return r
}
