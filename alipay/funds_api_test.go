package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestFundTransUniTransfer(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", "201806300011232301").
		Set("trans_amount", "0.01").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		SetBodyMap("payee_info", func(bm gopay.BodyMap) {
			bm.Set("identity", "85411418@qq.com")
			bm.Set("identity_type", "ALIPAY_LOGON_ID")
		})

	aliRsp, err := client.FundTransUniTransfer(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundAccountQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("alipay_user_id", "2088301409188095") /*.Set("account_type", "ACCTRANS_ACCOUNT")*/

	aliRsp, err := client.FundAccountQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundTransCommonQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("product_code", "TRANS_ACCOUNT_NO_PWD").
		Set("biz_scene", "DIRECT_TRANSFER").
		Set("order_id", "20190801110070000006380000250621")

	aliRsp, err := client.FundTransCommonQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundTransOrderQuery(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", "201806300011232301")

	aliRsp, err := client.FundTransOrderQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestFundAuthOrderAppFreeze(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_order_no", "8077735255938023").
		Set("out_request_no", "8077735255938032").
		Set("order_title", "预授权冻结").
		Set("amount", "0.01").
		Set("product_code", "PRE_AUTH_ONLINE")

	aliRsp, err := client.FundAuthOrderAppFreeze(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp:", aliRsp)
}

func TestClient_FundTransPagePay(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", "2018062800001").
		Set("trans_amount", "8.88").
		Set("product_code", "STD_APP_TRANSFER").
		Set("biz_scene", "PARTY_MEMBERSHIP_DUES").
		SetBodyMap("payee_info", func(b gopay.BodyMap) {
			b.Set("identity", "208812*****41234").
				Set("identity_type", "ALIPAY_USER_ID").
				Set("name", "黄龙国际有限公司")
		})

	aliRsp, err := client.FundTransPagePay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.FundTransPagePay(%+v),err:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestClient_FundTransAppPay(t *testing.T) {
	bm := make(gopay.BodyMap)
	bm.Set("out_biz_no", "20180628000035").
		Set("trans_amount", "8.88").
		Set("product_code", "STD_RED_PACKET").
		Set("biz_scene", "PERSONAL_PAY").
		Set("order_title", "钉钉拼手气红包")

	pageRedirectionData, err := client.FundTransAppPay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.FundTransAppPay(%+v),err:%+v", bm, err)
		return
	}
	xlog.Debug("pageRedirectionData: ", pageRedirectionData)
}
