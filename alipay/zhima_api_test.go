package alipay

import (
	"github.com/go-pay/gopay/pkg/util"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

// 芝麻企业信用信用评估初始化测试
func TestZhimaCreditEpSceneRatingInitialize(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_category", "ZMSCCO_5_1_1")
	bm.Set("product_code", "w1010100100000000001")
	bm.Set("out_order_no", "201805301527674106562F0000954216")
	bm.Set("user_id", "2088302248028263")

	aliRsp, err := client.ZhimaCreditEpSceneRatingInitialize(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneRatingInitialize(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 信用服务履约同步测试
func TestZhimaCreditEpSceneFulfillmentSync(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_order_no", "zme20181106154144733795615964647")
	bm.Set("out_order_no", "trade201805301527674106562F0000954217")
	bm.Set("biz_time", "2018-12-06 18:53:59")
	bm.Set("biz_ext_param", "{\"total_amount\":\"32890\"}")

	aliRsp, err := client.ZhimaCreditEpSceneFulfillmentSync(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneFulfillmentSync(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 加入信用服务测试
func TestZhimaCreditEpSceneAgreementUse(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("rating_order_no", "zme20181106154144730565715957902")
	bm.Set("out_order_no", "test201805301527674106562F0000954216")
	bm.Set("biz_time", "2018-12-06 18:53:59")
	bm.Set("provision_code", "P$ZMSCCO_5_1_1$00001")
	bm.Set("biz_ext_param", "{\"total_amount\":\"32890\"}")

	aliRsp, err := client.ZhimaCreditEpSceneAgreementUse(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneAgreementUse(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 取消信用服务测试
func TestZhimaCreditEpSceneAgreementCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_order_no", "zme20181106154144733795615964647")
	bm.Set("out_order_no", util.GetRandomString(64))
	bm.Set("biz_time", "2018-12-06 18:53:59")

	aliRsp, err := client.ZhimaCreditEpSceneAgreementCancel(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneAgreementCancel(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 信用服务履约同步(批量)测试
func TestZhimaCreditEpSceneFulfillmentlistSync(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_order_no", "zme20181106154144733795615964647")
	bm.SetBodyMap("fulfillment_info_list", func(bm gopay.BodyMap) {
		bm.Set("out_order_no", "trade201805301527674106562F0000954217")
		bm.Set("biz_time", "2018-12-06 18:53:59")
		bm.Set("biz_ext_param", "{\"total_amount\":\"32890\"}")
	})
	aliRsp, err := client.ZhimaCreditEpSceneFulfillmentlistSync(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditEpSceneFulfillmentlistSync(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 芝麻go用户数据回传测试
func TestZhimaCreditPeZmgoCumulationSync(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_no", "20195108518085620000")
	bm.Set("user_id", "2088602002015001")
	bm.Set("partner_id", "2088621805983504")
	bm.Set("out_biz_no", "2020081211223006150094012926289")
	bm.Set("biz_time", "2019-03-08 19:51:35")
	bm.Set("request_from", "ExternalMerchantSource")
	bm.Set("biz_action", "ORDER_PAID")
	// 可选
	bm.Set("cumulate_data_type", "AMOUNT")
	bm.Set("pay_out_biz_no", "2020081211223006150094012926289")
	bm.Set("has_alipay_trade", "false")
	bm.Set("ext_info", "{\"number\":\"20200303938282939\"}")
	bm.SetBodyMap("task_type_data", func(bm gopay.BodyMap) {
		bm.Set("name", "滴滴打车任务001")
	})
	bm.SetBodyMap("amount_type_data", func(bm gopay.BodyMap) {
		bm.Set("name", "优惠累计名称001")
		bm.SetBodyMap("trade_info", func(bm gopay.BodyMap) {
			bm.Set("trade_no", "2020081722001435461000061785")
			bm.Set("amount", "3.57")
		})
		bm.SetBodyMap("out_discount_infos", func(bm gopay.BodyMap) {
			bm.Set("discount_name", "滴滴打车优惠001")
			bm.Set("discount_type", "exclusiveBenefit")
			bm.Set("discount_amount", "8.75")
		})
	})
	aliRsp, err := client.ZhimaCreditPeZmgoCumulationSync(bm)
	if err != nil {
		xlog.Errorf("client.ZhimaCreditPeZmgoCumulationSync(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}
