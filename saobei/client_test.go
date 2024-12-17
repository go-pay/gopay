package saobei

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay/saobei/cert"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error
)

func TestMain(m *testing.M) {
	// 初始化通联客户端
	// instNo      string //商户系统机构号inst_no
	// key         string // 商户系统令牌
	// merchantNo  string // 支付系统：商户号
	// terminalId  string // 支付系统：商户号终端号
	// accessToken string // 支付系统： 令牌
	// isProd：是否是正式环境
	client, err = NewClient(cert.InstNo, cert.Key, cert.MerchantNo, cert.TerminalId, cert.AccessToken, false)
	if err != nil {
		xlog.Error(err)
		return
	}

	os.Exit(m.Run())
}
