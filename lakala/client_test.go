package lakala

import (
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

var (
	client         *Client
	err            error
	wxAppid        = "wxdaa2ab9ef87b5497"
	partnerCode    = "1368139502"
	credentialCode = "GFDS8j98rewnmgl45wHTt980jg543wmg"
)

func TestMain(m *testing.M) {
	// 初始化lakala户端
	//  wxAppid: 微信appid，微信通道要求必填
	//  PartnerCode: 商户编码，由4~6位大写字母或数字构成
	//  credentialCode: 系统为商户分配的开发校验码，请妥善保管，不要在公开场合泄露
	//  isProd: 是否生产环境
	client, err = NewClient(wxAppid, partnerCode, credentialCode, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	os.Exit(m.Run())
}
