package allinpay

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay/allinpay/cert"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
)

func TestMain(m *testing.M) {
	// 初始化通联客户端
	// cusId: 实际交易商户号
	// appid：平台分配的APPID
	// privateKey：商户的RSA私钥
	// publicKey：通联的公钥
	// isProd：是否是正式环境
	client, err = NewClient(cert.CusId, cert.AppId, cert.PrivateKey, cert.PublicKey, true)
	if err != nil {
		xlog.Error(err)
		return
	}

	os.Exit(m.Run())
}
