package douyin

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *Client
	err    error

	// 商户凭证：请填写自己的抖音支付商户信息
	Mchid      = ""
	SerialNo   = ""
	ApiKey     = "" // 32 字节接口加密密钥，用于回调 AES-256-GCM 解密
	PrivateKey = ``

	// 抖音支付平台证书（可注册多张，请填写自己的平台证书 PEM 与序列号）
	PlatformCert    = ``
	PlatformCertSNo = ""

	// 业务测试用参数：请按实际填写
	Appid     = ""
	NotifyUrl = "https://www.example.com/douyin/notify"
	Openid    = ""

	// 已存在的订单号（用于查询 / 关单 / 退款测试，请填写自己环境下的真实单号）
	TransactionId  = ""
	OutTradeNo     = ""
	OutRefundNo    = ""
	OutOrderNo     = ""
	OutReturnNo    = ""
	OutBillNo      = ""
	TransferBillNo = ""
)

func TestMain(m *testing.M) {
	// NewClient 初始化抖音支付客户端
	//	mchid：商户号
	//	serialNo：商户 API 证书序列号
	//	apiKey：接口加密密钥（32 字节）
	//	privateKey：商户 API 私钥 apiclient_key.pem 内容
	client, err = NewClient(Mchid, SerialNo, ApiKey, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return
	}

	// 注册抖音支付平台证书（用于响应/回调验签）
	if PlatformCert != "" && PlatformCertSNo != "" {
		if err = client.SetPlatformCert([]byte(PlatformCert), PlatformCertSNo); err != nil {
			xlog.Error(err)
			return
		}
	}

	// 打开 Debug 日志开关，输出请求/响应报文
	client.DebugSwitch = gopay.DebugOff

	// 给 HTTP 客户端设置整体超时，避免测试偶发卡住
	client.GetHttpClient().SetTimeout(15 * time.Second)

	os.Exit(m.Run())
}
