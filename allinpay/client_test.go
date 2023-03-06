package allinpay

import (
	"context"
	"github.com/go-pay/gopay/allinpay/cert"
	"github.com/go-pay/gopay/pkg/xlog"
	"os"
	"testing"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
	// 普通公钥模式时，验签使用
	//aliPayPublicKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
)

func TestMain(m *testing.M) {

	// 初始化支付宝客户端
	//    appid：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err = NewClient(cert.CusId, cert.AppId, cert.PrivateKey, cert.PublicKey)
	if err != nil {
		xlog.Error(err)
		return
	}

	os.Exit(m.Run())
}
