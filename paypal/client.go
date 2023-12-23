package paypal

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xhttp"
)

// Client PayPal支付客户端
type Client struct {
	Clientid       string
	Secret         string
	Appid          string
	AccessToken    string
	ExpiresIn      int
	IsProd         bool
	ctx            context.Context
	DebugSwitch    gopay.DebugSwitch
	hc             *xhttp.Client
	baseUrlProd    string
	baseUrlSandbox string
}

type Option func(*Client)

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool, options ...Option) (client *Client, err error) {
	if clientid == gopay.NULL || secret == gopay.NULL {
		return nil, gopay.MissPayPalInitParamErr
	}
	client = &Client{
		Clientid:       clientid,
		Secret:         secret,
		IsProd:         isProd,
		ctx:            context.Background(),
		DebugSwitch:    gopay.DebugOff,
		hc:             xhttp.NewClient(),
		baseUrlProd:    baseUrlProd,
		baseUrlSandbox: baseUrlSandbox,
	}
	for _, option := range options {
		option(client)
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 自动刷新Token
	go client.goAuthRefreshToken()
	return client, nil
}

// WithProxyUrl 设置代理 Url
func WithProxyUrl(proxyUrlProd, proxyUrlSandbox string) Option {
	return func(c *Client) {
		c.baseUrlProd = proxyUrlProd
		c.baseUrlSandbox = proxyUrlSandbox
	}
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetProxyUrl 设置代理 Url
// 使用场景：
// 1. 大陆直接调用 PayPal 接口响应较慢，可以在第三地例如硅谷部署代理服务器来加速请求
func (c *Client) SetProxyUrl(proxyUrlProd, proxyUrlSandbox string) {
	c.baseUrlProd = proxyUrlProd
	c.baseUrlSandbox = proxyUrlSandbox
}
