package alipay

import (
	"context"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay/cert"
	"github.com/go-pay/xlog"
)

var (
	ctx    = context.Background()
	client *ClientV3
	err    error
)

func TestMain(m *testing.M) {
	xlog.SetLevel(xlog.DebugLevel)
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err = NewClientV3(cert.Appid, cert.PrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// Debug开关，输出/关闭日志
	client.DebugSwitch = gopay.DebugOn

	// 设置biz_content加密KEY，设置此参数默认开启加密（目前不可用，设置后会报错）
	//client.SetAESKey("KvKUTqSVZX2fUgmxnFyMaQ==")

	// 传入证书内容
	err := client.SetCert(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	// 传入证书文件路径
	//err := client.SetCertSnByPath("cert/appPublicCert.crt", "cert/alipayRootCert.crt", "cert/alipayPublicCert.crt")
	if err != nil {
		xlog.Debug("SetCert:", err)
		return
	}
	os.Exit(m.Run())
}
