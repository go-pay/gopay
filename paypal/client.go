package paypal

import (
	"github.com/go-pay/gopay"
)

// Client PayPal支付客
type Client struct {
	DebugSwitch gopay.DebugSwitch
}

// NewClient 初始化PayPal支付客户端
func NewClient() (client *Client, err error) {

	client = &Client{
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}
