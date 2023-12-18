package alipay

import (
	"testing"

	"github.com/go-pay/gopay"
	"github.com/go-pay/xlog"
)

func TestKoubeiTradeOrderPrecreate(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "20181120111040030100030100002400")
	bm.Set("biz_type", "POST_ORDER_PAY")

	aliRsp, err := client.KoubeiTradeOrderPrecreate(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeOrderPrecreate(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeItemorderBuy(t *testing.T) {
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

	aliRsp, err := client.KoubeiTradeItemorderBuy(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeItemorderBuy(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeOrderConsult(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "0ad1e47b1500473065347103327127")
	bm.Set("user_id", "2088212151390950")
	bm.Set("total_amount", "88.88")
	bm.Set("shop_id", "2015051100077000000000000300")

	aliRsp, err := client.KoubeiTradeOrderConsult(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeOrderConsult(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeItemorderRefund(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("order_no", "20171114111040030100030100002400")
	bm.Set("out_request_no", "B201701180000002")
	bm.Set("total_amount", "88.88")
	bm.SetBodyMap("refund_infos", func(bm gopay.BodyMap) {
		bm.Set("item_order_no", "201701220000008000000001")
		bm.Set("amount", "10.00")
	})

	aliRsp, err := client.KoubeiTradeItemorderRefund(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeItemorderRefund(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeItemorderQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("order_no", "20160831001040010900010000000001")
	bm.Set("out_request_no", "20160831001040010900010000000001")
	bm.Set("partner_id", "20180000000001")
	bm.Set("status", "CLOSE")
	bm.Set("buyer_id", "2088102164040745")
	bm.Set("biz_product", "ONLINE_TRADE_PAY")
	bm.Set("gmt_create", "2016-09-29 00:00:00")
	bm.Set("gmt_modified", "2016-09-29 00:00:00")
	bm.Set("total_amount", "25.00")
	bm.Set("real_pay_amount", "20.00")
	bm.Set("deliver_seller_real_amount", "5.00")
	bm.SetBodyMap("item_order_vo", func(bm gopay.BodyMap) {
		bm.Set("item_order_no", "2015060400076000000000000000")
		bm.Set("sku_id", "2015060400076000000000012100")
		bm.Set("quantity", "2")
		bm.Set("price", "10.00")
		bm.Set("status", "SUCCESS")
	})

	aliRsp, err := client.KoubeiTradeItemorderQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeItemorderQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeTicketTicketcodeSend(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "20161029032144700000000001467823")
	bm.Set("send_order_no", "20170428001640010400000000006800")
	bm.Set("send_token", "eed3206686d349d2a22c3d9ada35c0f8")
	bm.Set("order_no", "20180404111040030100130500594477")
	bm.SetBodyMap("isv_ma_list", func(bm gopay.BodyMap) {
		bm.Set("code", "GAL358309629")
		bm.Set("num", "2")
	})

	aliRsp, err := client.KoubeiTradeTicketTicketcodeSend(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeTicketTicketcodeSend(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeTicketTicketcodeDelay(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "20161029032144700000000001467823")
	bm.Set("end_date", "2017-09-09 12:12:12")
	bm.Set("ticket_code", "358309629059")
	bm.Set("code_type", "INTERNAL_CODE")
	bm.Set("order_no", "20180404111040030100130500594477")

	aliRsp, err := client.KoubeiTradeTicketTicketcodeDelay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeTicketTicketcodeDelay(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeTicketTicketcodeQuery(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("ticket_code", "016569843362")
	bm.Set("shop_id", "2017071200077000000039734370")

	aliRsp, err := client.KoubeiTradeTicketTicketcodeQuery(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeTicketTicketcodeQuery(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}

func TestKoubeiTradeTicketTicketcodeCancel(t *testing.T) {
	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("request_id", "2016102903214476899999999")
	bm.Set("request_biz_no", "2016102903214476899999999")
	bm.Set("ticket_code", "016569843362")

	aliRsp, err := client.KoubeiTradeTicketTicketcodeCancel(ctx, bm)
	if err != nil {
		xlog.Errorf("client.KoubeiTradeTicketTicketcodeCancel(%+v),error:%+v", bm, err)
		return
	}
	xlog.Debug("aliRsp:", *aliRsp)
}
