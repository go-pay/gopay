package apple

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/util"
)

// Client AppleClient
type Client struct {
	IssuerId        string // 从 AppStore Connect 获得的 Issuer Id
	BundleId        string // app包名
	AppleKeyId      string // 内购密钥id
	ApplePrivateKey string // 内购密钥.p8文件内容
}

// NewClient 初始化Apple客户端
// mchid：商户ID 或者服务商模式的 sp_mchid
// serialNo：商户API证书的证书序列号
// apiV3Key：APIv3Key，商户平台获取
// privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
func NewClient(mchid, serialNo, apiV3Key, privateKey string) (client *Client, err error) {
	if mchid == util.NULL || serialNo == util.NULL || apiV3Key == util.NULL || privateKey == util.NULL {
		return nil, gopay.MissWechatInitParamErr
	}

	client = &Client{}
	return client, nil
}
