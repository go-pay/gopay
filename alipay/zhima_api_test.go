package alipay

import (
	"errors"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/util"
	"github.com/go-pay/xlog"
)

// 芝麻企业信用信用评估初始化测试
func TestZhimaCreditEpSceneRatingInitialize(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_category", "ZMSCCO_5_1_1")
	bm.Set("product_code", "w1010100100000000001")
	bm.Set("out_order_no", "201805301527674106562F0000954216")
	bm.Set("user_id", "2088302248028263")

	aliRsp, err := client.ZhimaCreditEpSceneRatingInitialize(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditEpSceneRatingInitialize(%+v),error:%+v", bm, err)
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

	aliRsp, err := client.ZhimaCreditEpSceneFulfillmentSync(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditEpSceneFulfillmentSync(%+v),error:%+v", bm, err)
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

	aliRsp, err := client.ZhimaCreditEpSceneAgreementUse(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditEpSceneAgreementUse(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

// 取消信用服务测试
func TestZhimaCreditEpSceneAgreementCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("credit_order_no", "zme20181106154144733795615964647")
	bm.Set("out_order_no", util.RandomString(64))
	bm.Set("biz_time", "2018-12-06 18:53:59")

	aliRsp, err := client.ZhimaCreditEpSceneAgreementCancel(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditEpSceneAgreementCancel(%+v),error:%+v", bm, err)
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
	aliRsp, err := client.ZhimaCreditEpSceneFulfillmentlistSync(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditEpSceneFulfillmentlistSync(%+v),error:%+v", bm, err)
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
	aliRsp, err := client.ZhimaCreditPeZmgoCumulationSync(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoCumulationSync(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 商家芝麻GO累计数据回传接口测试
func TestZhimaMerchantZmgoCumulateSync(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "20195108518085620000")
	bm.Set("user_id", "2088602002015001")
	bm.Set("provider_pid", "2088621805983504")
	bm.Set("out_biz_no", "2020081211223006150094012926289")
	bm.Set("biz_time", "2019-03-08 19:51:35")
	bm.Set("biz_action", "POSITIVE")
	bm.Set("sub_biz_action", "ADD")
	bm.Set("data_type", "TASK")
	// 可选
	bm.Set("refer_out_biz_no", "2020081211223006150094012926289")
	bm.SetBodyMap("amount_type_sync_data", func(bm gopay.BodyMap) {
		bm.Set("task_desc", "完成一次任务001")
		bm.Set("task_amount", "3.57")
		bm.Set("trade_no", "2020081722001435461000061785")
		bm.Set("has_alipay_trade", "false")
		bm.Set("discount_desc", "消费满减优惠001")
		bm.Set("discount_amount", "8.75")
	})
	bm.SetBodyMap("times_type_sync_data", func(bm gopay.BodyMap) {
		bm.Set("task_desc", "完成一次任务001")
		bm.Set("task_times", "1")
		bm.Set("task_amount", "3.57")
		bm.Set("discount_desc", "消费满减优惠001")
		bm.Set("discount_amount", "8.75")
	})
	bm.SetBodyMap("discount_type_sync_data", func(bm gopay.BodyMap) {
		bm.Set("discount_desc", "消费满减优惠001")
		bm.Set("discount_amount", "8.75")
	})

	aliRsp, err := client.ZhimaMerchantZmgoCumulateSync(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaMerchantZmgoCumulateSync(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 商家芝麻GO累计数据查询接口测试
func TestZhimaMerchantZmgoCumulateQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "20195108518085620000")
	bm.Set("user_id", "2088602002015001")
	bm.Set("provider_pid", "2088621805983504")
	// 可选
	bm.Set("need_detail", "false")
	bm.Set("page_no", "1")
	bm.Set("page_size", "20")

	aliRsp, err := client.ZhimaMerchantZmgoCumulateQuery(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaMerchantZmgoCumulateQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO签约关单测试
func TestZhimaCreditPeZmgoBizoptClose(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "2088302841345600")
	bm.Set("partner_id", "2088302424614288")
	bm.Set("out_request_no", "99202005050100930053707258")
	bm.Set("template_id", "2021012300020903090008858258")

	aliRsp, err := client.ZhimaCreditPeZmgoBizoptClose(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoBizoptClose(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO结算退款接口测试
func TestZhimaCreditPeZmgoSettleRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "ZMGO_AGR2021010510020604080000006001")
	bm.Set("partner_id", "2088302424614288")
	bm.Set("alipay_user_id", "2088302841345600")
	bm.Set("refund_amount", "3.00")
	bm.Set("out_request_no", "99202005050100930053707258")
	// 可选
	bm.Set("memo", "退款")
	bm.Set("withhold_plan_no", "ZMGO_WHD2021010510020603410000006001")
	bm.Set("refund_type", "MEMBER_FEE_REFUND")

	aliRsp, err := client.ZhimaCreditPeZmgoSettleRefund(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoSettleRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO签约预创单测试
func TestZhimaCreditPeZmgoPreorderCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("partner_id", "2088437463829741")
	bm.Set("template_id", "2020122200020903760008552025")
	bm.Set("out_request_no", "ORDER_12345678")
	bm.Set("biz_time", "2016-07-06 00:00:02")
	// 可选
	bm.Set("freeze_amount", "1.00")
	bm.Set("timeout_express", "15m")
	bm.Set("alipay_user_id", "2088759402857364")
	bm.Set("partner_user_identifier", "user102934889234")
	bm.Set("isv_pid", "2088374762857463")
	bm.Set("sign_aisle_data", "业务方签约标识")
	bm.Set("expire_aisle_data", "业务方签约标识")
	bm.Set("pay_aisle_data", "业务方签约标识")
	bm.SetBodyMap("ext_template_conf", func(bm gopay.BodyMap) {
		bm.Set("xxhm_info_id", "687542")
		bm.Set("buyer_id", "11212321121")
	})

	aliRsp, err := client.ZhimaCreditPeZmgoPreorderCreate(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoPreorderCreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO协议解约测试
func TestZhimaCreditPeZmgoAgreementUnsign(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "ZMGO_AGR2020122710020604120000000001")
	bm.Set("partner_id", "2088302424614288")
	// 可选
	bm.Set("alipay_user_id", "2088302841345600")
	bm.Set("quit_type", "SETTLE_APPLY_QUIT")

	aliRsp, err := client.ZhimaCreditPeZmgoAgreementUnsign(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoAgreementUnsign(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻Go协议查询接口测试
func TestZhimaCreditPeZmgoAgreementQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "20185513447859192007")
	bm.Set("alipay_user_id", "2088101117955611")

	aliRsp, err := client.ZhimaCreditPeZmgoAgreementQuery(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoAgreementQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻Go解冻接口测试
func TestZhimaCreditPeZmgoSettleUnfreeze(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("agreement_id", "20185513447859192007")
	bm.Set("out_request_no", "8077735255938032")
	bm.Set("unfreeze_amount", "3.00")
	bm.Set("biz_time", "2014-09-15 11:23:04")
	bm.Set("alipay_user_id", "2088101117955611")
	// 可选
	bm.Set("partner_id", "2088411663864410")
	bm.Set("order_title", "2088411663864410")
	bm.SetBodyMap("unfreeze_extend_params", func(bm gopay.BodyMap) {
		bm.Set("total_real_pay_amount", "3.00")
		bm.Set("total_discount_amount", "3.00")
		bm.Set("total_task_count", "0")
		bm.Set("quit_type", "SETTLE_APPLY_QUIT")
	})

	aliRsp, err := client.ZhimaCreditPeZmgoSettleUnfreeze(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoSettleUnfreeze(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO支付下单链路签约申请测试
func TestZhimaCreditPeZmgoPaysignApply(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "2088302841345600")
	bm.Set("partner_id", "2088302424614288")
	bm.Set("template_id", "2021012300020903090008858258")
	bm.Set("merchant_app_id", "2021001118641054")
	bm.Set("out_request_no", "99202005050100930053707258")
	bm.Set("biz_time", "2016-07-06 00:00:02")
	bm.Set("timeout_express", "1m")
	// 可选

	aliRsp, err := client.ZhimaCreditPeZmgoPaysignApply(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoPaysignApply(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 芝麻GO支付下单链路签约确认测试
func TestZhimaCreditPeZmgoPaysignConfirm(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "20195108518085620000")
	bm.Set("partner_id", "2088302424614288")
	bm.Set("merchant_app_id", "2021001118641054")
	bm.Set("zmgo_opt_no", "ZMGO_OPT2021040110020607570001198871")
	bm.Set("biz_type", "hongbaoqiandao")
	// 可选

	aliRsp, err := client.ZhimaCreditPeZmgoPaysignConfirm(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCreditPeZmgoPaysignConfirm(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 职得工作证信息匹配度查询测试
func TestZhimaCustomerJobworthAdapterQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 可选
	bm.Set("user_id", "2088302887413828")
	bm.Set("cert_no", "01011219900101XXXX")
	bm.Set("user_name", "张三")
	bm.Set("cert_type", "1")
	bm.Set("job_name", "java开发")
	bm.SetBodyMap("adapter", func(bm gopay.BodyMap) {
		bm.Set("edu_level", "BACHELOR")
		bm.Set("skill_certificate", "电工证")
		bm.Set("age", "18.35")
		bm.Set("gender", "男")
		bm.Set("city", "330100")
		bm.Set("recommend", "5")
	})

	aliRsp, err := client.ZhimaCustomerJobworthAdapterQuery(ctx, bm)
	if err != nil {
		//xlog.Errorf("client.ZhimaCustomerJobworthAdapterQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}

// 职得工作证外部渠道应用数据回流测试
func TestZhimaCustomerJobworthSceneUse(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	// 可选
	bm.Set("scene_type", "zhima_job_ext_exposure_info")
	bm.SetBodyMap("extra_info", func(bm gopay.BodyMap) {
		bm.Set("user_id", "2088302887413828")
		bm.Set("cert_no", "01011219900101XXXX")
		bm.Set("user_name", "张三")
		bm.Set("job_name", "java开发")
		bm.Set("job_id", "123456")
		bm.Set("scene_time", "1624278367975")
		bm.Set("job_category_id", "project")
		bm.Set("company_name", "58")
		bm.Set("company_id", "123456789")
		bm.Set("job_category", "技术")
		bm.Set("employer_visit", "true")
		bm.Set("self_visit", "true")
	})

	aliRsp, err := client.ZhimaCustomerJobworthSceneUse(ctx, bm)
	if err != nil {
		if errors.Is(err, gopay.UnmarshalErr) {
			//xlog.Errorf("%v", err)
			return
		}
		xlog.Errorf("client.ZhimaCustomerJobworthSceneUse(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)

}
