package bytedance

import (
	"context"

	"github.com/go-pay/gopay"
)

// Client 字节担保支付
type Client struct {
	Token       string
	Mchid       string
	Salt        string
	ctx         context.Context
	DebugSwitch gopay.DebugSwitch
}

// NewClient 初始化字节客户端
// 	token：Token 令牌
//	mchid：商户号
//	salt：SALT
func NewClient(token, mchid, salt string) (client *Client, err error) {
	client = &Client{
		Token:       token,
		Mchid:       mchid,
		Salt:        salt,
		ctx:         context.Background(),
		DebugSwitch: gopay.DebugOff,
	}
	return client, nil
}
