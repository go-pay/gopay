package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
)

func TestTradeOrderPrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "20181120111040030100030100002400")
	bm.Set("biz_type", "POST_ORDER_PAY")

	aliRsp, err := client.KoubeiTradeOrderPrecreate(bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeOrderPrecreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestTradeItemorderBuy(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_order_no", "A2018011200000001")
	bm.Set("subject", "星巴克礼品卡")
	bm.Set("biz_product", "ONLINE_PURCHASE")
	bm.Set("biz_scene", "giftCard")
	bm.Set("shop_id", "2015051100077000000000000300")
	bm.Set("buyer_id", "2088102015433735")
	bm.Set("total_amount", "100.00")
	bm.SetBodyMap("item_order_details", func(bm gopay.BodyMap) {
		bm.Set("sku_id", "2015060400076000000000012100")
		bm.Set("original_price", "50.00")
		bm.Set("price", "10.00")
		bm.Set("quantity", "10")
	})

	aliRsp, err := client.KoubeiTradeItemorderBuy(bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeItemorderBuy(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
