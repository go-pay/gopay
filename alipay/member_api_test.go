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
	merchant["return_url"] = "https://www.fmm.ink"
	bm.Set("merchant_config", merchant)

	// 发起请求
	aliRsp, err := client.UserCertifyOpenInit(bm)
	if err != nil {
		xlog.Error(err)
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
		xlog.Error(err)
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
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response.Passed:", aliRsp.Response.Passed)
}

func TestUserAgreementExecutionplanModify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_no", "20185909000458725113")
	bm.Set("deduct_time", "2019-05-12")
	bm.Set("memo", "用户已购买半年包，需延期扣款时间")

	// 发起请求
	aliRsp, err := client.UserAgreementExecutionplanModify(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAgreementTransfer(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_no", "20170322450983769228")
	bm.Set("target_product_code", "CYCLE_PAY_AUTH_P")
	bm.SetBodyMap("period_rule_params", func(bm gopay.BodyMap) {
		bm.Set("period_type", "DAY")
		bm.Set("period", "3")
		bm.Set("execute_time", "20190-01-23")
		bm.Set("single_amount", "10.99")
	})

	// 发起请求
	aliRsp, err := client.UserAgreementTransfer(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserTwostageCommonUse(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("dynamic_id", "286861260475412123")
	bm.Set("sence_no", "20170718xxxxxxxx")
	bm.Set("pay_pid", "2088702093900999")

	// 发起请求
	aliRsp, err := client.UserTwostageCommonUse(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAuthZhimaorgIdentityApply(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("cert_type", "NATIONAL_LEGAL")
	bm.Set("cert_no", "330701199901011311")
	bm.Set("name", "中国移动有限公司")

	// 发起请求
	aliRsp, err := client.UserAuthZhimaorgIdentityApply(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserCharityRecordexistQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("partner_id", "2088111122223333")
	bm.Set("user_id", "2088111122223333")

	// 发起请求
	aliRsp, err := client.UserCharityRecordexistQuery(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserAlipaypointSend(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("budget_code", "02559A591572")
	bm.Set("partner_biz_no", "011022222222212")
	bm.Set("point_amount", "1")

	// 发起请求
	aliRsp, err := client.UserAlipaypointSend(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestMemberDataIsvCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("member_card_id", "2014323100009")
	bm.Set("member_source", "alipay")
	bm.Set("member_status", "1")
	bm.Set("gmt_merber_card_create", "2017-02-17 20:11:54")
	bm.Set("parter_id", "2088902248579233")

	// 发起请求
	aliRsp, err := client.MemberDataIsvCreate(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestUserFamilyArchiveQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("archive_token", "2020050200286001170017000004861")

	// 发起请求
	aliRsp, err := client.UserFamilyArchiveQuery(bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
