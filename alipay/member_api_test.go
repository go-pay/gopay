package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func TestClient_UserCertifyOpenInit(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("outer_order_no", "ZGYD201809132323000001234")
	// 认证场景码：FACE：多因子人脸认证，CERT_PHOTO：多因子证照认证，CERT_PHOTO_FACE ：多因子证照和人脸认证，SMART_FACE：多因子快捷认证
	bm.Set("biz_code", "FACE")
	// 需要验证的身份信息参数，格式为json
	identity := make(map[string]string)
	identity["identity_type"] = "CERT_INFO"
	identity["cert_type"] = "IDENTITY_CARD"
	identity["cert_name"] = "张三"
	identity["cert_no"] = "310123199012301234"
	bm.Set("identity_param", identity)
	// 商户个性化配置，格式为json
	merchant := make(map[string]string)
	merchant["return_url"] = "https://www.fumm.cc"
	bm.Set("merchant_config", merchant)

	// 发起请求
	aliRsp, err := client.UserCertifyOpenInit(bm)
	if err != nil {
		xlog.Errorf("client.UserCertifyOpenInit(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_UserCertifyOpenCertify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "53827f9d085b3ce43938c6e5915b4729")

	// 发起请求
	certifyUrl, err := client.UserCertifyOpenCertify(bm)
	if err != nil {
		xlog.Errorf("client.UserCertifyOpenCertify(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("certifyUrl:", certifyUrl)
}

func TestClient_UserCertifyOpenQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
	bm.Set("certify_id", "OC201809253000000393900404029253")

	// 发起请求
	aliRsp, err := client.UserCertifyOpenQuery(bm)
	if err != nil {
		xlog.Errorf("client.UserCertifyOpenQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response.Passed:", aliRsp.Response.Passed)
}
