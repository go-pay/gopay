package paypal

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
	"github.com/go-pay/gopay/pkg/xhttp"
)

// Client PayPal支付客户端
type Client struct {
	Clientid    string
	Secret      string
	Appid       string
	AccessToken string
	ExpiresIn   int
	IsProd      bool
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
	hc          *xhttp.Client
}

// NewClient 初始化PayPal支付客户端
func NewClient(clientid, secret string, isProd bool) (client *Client, err error) {
	if clientid == util.NULL || secret == util.NULL {
		return nil, gopay.MissPayPalInitParamErr
	}
	client = &Client{
		Clientid:    clientid,
		Secret:      secret,
		IsProd:      isProd,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
		hc:          xhttp.NewClient(),
	}
	_, err = client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 自动刷新Token
	go client.goAuthRefreshToken()
	return client, nil
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}
