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

// NewClient , default tls.Config{InsecureSkipVerify: true}
func NewClient() (client *Client) {
	client = &Client{
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
	return client
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

func (c *Client) Req(typeStr ...RequestType) *Request {
	var (
		reqTp = TypeJSON // default
		resTp = TypeJSON // default
		tLen  = len(typeStr)
	)
	switch {
	case tLen == 1:
		tpp := typeStr[0]
		if _, ok := types[tpp]; ok {
			reqTp = tpp
			// default resTp = reqTp
			resTp = tpp
		}
	case tLen > 1:
		// first param is request type
		tpp := typeStr[0]
		if _, ok := types[tpp]; ok {
			reqTp = tpp
		}
		// second param is response data type
		stpp := typeStr[1]
		if _, ok := types[stpp]; ok {
			resTp = stpp
		}
	}
	r := &Request{
		client:        c,
		Header:        make(http.Header),
		requestType:   reqTp,
		unmarshalType: string(resTp),
	}
	r.Header.Set("Content-Type", types[reqTp])
	return r
}
