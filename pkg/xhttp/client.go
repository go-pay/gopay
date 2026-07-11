package xhttp

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
	bodySizeMB int // body size limit (MB), default is 10MB
}

func defaultClient() *Client {
	tp := http.DefaultTransport.(*http.Transport).Clone()
	tp.MaxIdleConns = 5000
	tp.MaxIdleConnsPerHost = 1000
	tp.MaxConnsPerHost = 3000
	return &Client{
		HttpClient: &http.Client{
			Timeout:   60 * time.Second,
			Transport: tp,
		},
		bodySizeMB: 10,
	}
}

// NewClient 默认启用 TLS 证书校验。
// 沙箱或自签证书等场景需跳过校验时，调用 SetHttpTLSConfig 传入自定义 *tls.Config。
func NewClient() (client *Client) {
	return defaultClient()
}

func (c *Client) SetTransport(transport http.RoundTripper) (client *Client) {
	c.HttpClient.Transport = transport
	return c
}

// SetHttpTransport 等价于 SetTransport，仅为兼容旧 API 保留。
// 新代码请直接用 SetTransport。
func (c *Client) SetHttpTransport(transport *http.Transport) (client *Client) {
	return c.SetTransport(transport)
}

// SetHttpTLSConfig 在底层 *http.Transport 上设置 TLS 配置。
// 如果先调用 SetTransport 把 transport 替换成了非 *http.Transport 的实现，
// 调用本方法将 panic —— 而不是静默丢失 TLS 配置导致请求以默认 TLS 设置出网。
func (c *Client) SetHttpTLSConfig(tlsCfg *tls.Config) (client *Client) {
	ht, ok := c.HttpClient.Transport.(*http.Transport)
	if !ok {
		panic("xhttp: SetHttpTLSConfig requires the underlying Transport to be *http.Transport; current transport has been replaced via SetTransport")
	}
	ht.TLSClientConfig = tlsCfg
	return c
}

func (c *Client) SetTimeout(timeout time.Duration) (client *Client) {
	c.HttpClient.Timeout = timeout
	return c
}

// SetBodySize 设置响应体大小上限（MB），默认 10MB。
func (c *Client) SetBodySize(sizeMB int) (client *Client) {
	c.bodySizeMB = sizeMB
	return c
}

// Req 构造一个新的 Request。
// typeStr 第一个参数为请求 Content-Type，第二个参数为响应解析类型。
// 默认请求体类型 TypeJSON，响应类型 ResTypeJSON。
//
// 调用方必须先通过 NewClient() 构造 *Client；对 nil 接收者调用 Req 将 panic。
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
	r := &Request{
		client:       c,
		Header:       make(http.Header),
		requestType:  reqTp,
		responseType: resTp,
	}
	r.Header.Set("Content-Type", _ReqContentTypeMap[reqTp])
	return r
}
