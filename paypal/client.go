package paypal

import (
	"context"
	"github.com/go-pay/gopay/pkg/xlog"
	"time"

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
	token, err := client.GetAccessToken()
	if err != nil {
		return nil, err
	}
	// 在到期结束前的一半时间内，自动刷新token
	go func(token *AccessToken) {
		ticker := time.NewTicker(time.Duration(token.ExpiresIn / 2))
		for {
			select {
			case <-ticker.C:
				tokenNew, err := client.GetAccessToken()
				if err != nil {
					xlog.Errorf("PayPal GetAccessToken Error: %s", err.Error())
					continue
				}
				client.AccessToken = tokenNew.AccessToken
			}
		}
	}(token)

	return client, nil
}

// SetBodySize 设置http response body size(MB)
func (c *Client) SetBodySize(sizeMB int) {
	if sizeMB > 0 {
		c.hc.SetBodySize(sizeMB)
	}
}
