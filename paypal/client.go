package paypal

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xhttp"
	"github.com/go-pay/xlog"
)

// Client PayPal支付客户端
type Client struct {
	Clientid         string
	Secret           string
	Appid            string
	AccessToken      string
	ExpiresIn        int
	IsProd           bool
	ctx              context.Context
	DebugSwitch      gopay.DebugSwitch
	logger           xlog.XLogger
	hc               *xhttp.Client
	baseUrlProd      string
	baseUrlSandbox   string
	autoRefreshToken bool
	headerKeyMap     map[string]string
}

type Option func(*Client)

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool, options ...Option) (client *Client, err error) {
	if clientid == gopay.NULL || secret == gopay.NULL {
		return nil, gopay.MissPayPalInitParamErr
	}
	logger := xlog.NewLogger()
	logger.SetLevel(xlog.DebugLevel)
	client = &Client{
		Clientid:         clientid,
		Secret:           secret,
		IsProd:           isProd,
		ctx:              context.Background(),
		DebugSwitch:      gopay.DebugOff,
		logger:           logger,
		hc:               xhttp.NewClient(),
		baseUrlProd:      baseUrlProd,
		baseUrlSandbox:   baseUrlSandbox,
		autoRefreshToken: true,
		headerKeyMap:     make(map[string]string),
	}
	for _, option := range options {
		option(client)
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 自动刷新Token
	if client.autoRefreshToken {
		go client.goAuthRefreshToken()
	}
	return client, nil
}

// WithProxyUrl 设置代理 Url
func WithProxyUrl(proxyUrlProd, proxyUrlSandbox string) Option {
	return func(c *Client) {
		c.baseUrlProd = proxyUrlProd
		c.baseUrlSandbox = proxyUrlSandbox
	}
}

// WithHttpClient 设置自定义的xhttp.Client
func WithHttpClient(client *xhttp.Client) Option {
	return func(c *Client) {
		c.hc = client
	}
}

// WithoutAutoRefreshToken 设置不自动刷新Token
func WithoutAutoRefreshToken() Option {
	return func(c *Client) {
		c.autoRefreshToken = false
	}
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}

// SetHttpClient 设置自定义的xhttp.Client
func (c *Client) SetHttpClient(client *xhttp.Client) {
	if client != nil {
		c.hc = client
	}
}

func (c *Client) SetLogger(logger xlog.XLogger) {
	if logger != nil {
		c.logger = logger
	}
}

// SetProxyUrl 设置代理 Url
// 使用场景：
// 1. 大陆直接调用 PayPal 接口响应较慢，可以在第三地例如硅谷部署代理服务器来加速请求
func (c *Client) SetProxyUrl(proxyUrlProd, proxyUrlSandbox string) {
	c.baseUrlProd = proxyUrlProd
	c.baseUrlSandbox = proxyUrlSandbox
}

// SetRequestHeader 设置自定义的header
// defaultVal: 默认值
func (c *Client) SetRequestHeader(key string, defaultVal ...string) {
	if key != "" {
		if len(defaultVal) > 0 {
			c.headerKeyMap[key] = defaultVal[0]
		} else {
			c.headerKeyMap[key] = ""
		}
	}
}

// CleanRequestHeader 清理自定义的header
func (c *Client) CleanRequestHeader() {
	c.headerKeyMap = make(map[string]string)
}
