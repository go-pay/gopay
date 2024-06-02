package alipay

import (
	"io"
	"os"
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestOpenAppQrcodeCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("url_param", "page/component/component-pages/view/view").
		Set("query_param", "x=1").
		Set("describe", "二维码描述")

	// 发起请求
	aliRsp, err := client.OpenAppQrcodeCreate(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingCampaignCashCreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("coupon_name", "test_name").
		Set("prize_type", "0").
		Set("total_money", "0.1").
		Set("total_num", "1").
		Set("prize_msg", "test_prize_msg").
		Set("start_time", "0").
		Set("end_time", "0").
		Set("merchant_link", "0")

	// 发起请求
	aliRsp, err := client.MarketingCampaignCashCreate(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingCampaignCashTrigger(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("user_id", "1").
		Set("crowd_no", "1").
		Set("login_id", "1").
		Set("order_price", "0").
		Set("out_biz_no", "1")

	// 发起请求
	aliRsp, err := client.MarketingCampaignCashTrigger(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingCampaignCashStatusModify(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("crowd_no", "1").
		Set("camp_status", "1")
	// 发起请求
	aliRsp, err := client.MarketingCampaignCashStatusModify(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingCampaignCashListQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("camp_status", "1").
		Set("page_size", "1").
		Set("page_index", "1")
	// 发起请求
	aliRsp, err := client.MarketingCampaignCashListQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingCampaignCashDetailQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("crowd_no", "POYb84lfiKVdIfERAYsqPL_KQRIpfQbl47xfRmmPBlDMnSZ96O-zxUfKlHp5cxmx")

	// 发起请求
	aliRsp, err := client.MarketingCampaignCashDetailQuery(ctx, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("aliRsp.Response:", aliRsp.Response)
}

func TestMarketingMaterialImageUpload(t *testing.T) {
	// 请求参数
	logo, err := os.Open("../logo.png")
	if err != nil {
		xlog.Errorf("os.Open(%s),error:%+v", "../logo.png", err)
		return
	}
	xlog.Warnf("fileName: %s", logo.Name())
	allBs, err := io.ReadAll(logo)
	if err != nil {
		xlog.Errorf("io.ReadAll(%s),error:%+v", logo.Name(), err)
		return
	}
	f := &gopay.File{
		Name:    "logo.png",
		Content: allBs,
	}
	bm := make(gopay.BodyMap)
	bm.Set("file_key", "PROMO_VOUCHER_IMAGE").
		SetFormFile("file_content", f)
	aliRsp, err := client.MarketingMaterialImageUpload(ctx, bm)
	if err != nil {
		xlog.Errorf("client.MarketingMaterialImageUpload(),error:%+v", err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
