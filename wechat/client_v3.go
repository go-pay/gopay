package wechat

import (
	"sync"

	"github.com/iGoogle-ink/gopay"
)

type ClientV3 struct {
	AppId       string
	MchId       string
	ApiKey      string
	BaseURL     string
	IsProd      bool
	DebugSwitch gopay.DebugSwitch
	mu          sync.RWMutex
}

// 初始化微信客户端 V3
//	appId：应用ID
//	mchId：商户ID
//	ApiKey：API秘钥值
//	IsProd：是否是正式环境
func NewClientV3(appId, mchId, apiKey string, isProd bool) (client *Client) {
	return &Client{
		AppId:       appId,
		MchId:       mchId,
		ApiKey:      apiKey,
		IsProd:      isProd,
		DebugSwitch: gopay.DebugOff,
	}
}
