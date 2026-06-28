package lakala

import (
	"os"
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

var (
	client         *Client
	err            error
	partnerCode    = "xxxxxxx"
	credentialCode = "xxxxxxx"
)

func TestMain(m *testing.M) {
	// 初始化lakala户端
	//  PartnerCode: 商户编码，由4~6位大写字母或数字构成
	//  credentialCode: 系统为商户分配的开发校验码，请妥善保管，不要在公开场合泄露
	//  isProd: 是否生产环境
	client, err = NewClient(partnerCode, credentialCode, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn
	// 给 HTTP 客户端设置整体超时，避免支付宝某些接口偶发卡住导致 go test 整体超时
	client.GetHttpClient().SetTimeout(15 * time.Second)

	os.Exit(m.Run())
}
