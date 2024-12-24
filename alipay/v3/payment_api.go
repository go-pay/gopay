package alipay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-pay/gopay"
)

// 统一收单交易支付接口 alipay.trade.pay
// StatusCode = 200 is success
func (a *ClientV3) TradePay(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePayRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject", "auth_code", "scene")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradePay, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradePay, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradePayRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易查询 alipay.trade.query
// StatusCode = 200 is success
func (a *ClientV3) TradeQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeQueryRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易退款接口 alipay.trade.refund
// StatusCode = 200 is success
func (a *ClientV3) TradeRefund(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRefundRsp, err error) {
	err = bm.CheckEmptyError("refund_amount")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeRefund, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeRefund, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeRefundRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易退款查询 alipay.trade.fastpay.refund.query
// StatusCode = 200 is success
func (a *ClientV3) TradeFastPayRefundQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeFastPayRefundQueryRsp, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeFastPayRefundQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeFastPayRefundQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeFastPayRefundQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易撤销接口 alipay.trade.cancel
// StatusCode = 200 is success
func (a *ClientV3) TradeCancel(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCancelRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeCancel, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeCancel, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCancelRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易关闭接口 alipay.trade.close
// StatusCode = 200 is success
func (a *ClientV3) TradeClose(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCloseRsp, err error) {
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("trade_no") == gopay.NULL {
		return nil, errors.New("out_trade_no and trade_no are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeClose, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeClose, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCloseRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 查询对账单下载地址 alipay.data.dataservice.bill.downloadurl.query
// StatusCode = 200 is success
func (a *ClientV3) DataBillDownloadUrlQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *DataBillDownloadUrlQueryRsp, err error) {
	err = bm.CheckEmptyError("bill_type", "bill_date")
	if err != nil {
		return nil, err
	}
	uri := v3DataBillDownloadUrlQuery + "?" + bm.EncodeURLParams()
	authorization, err := a.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doGet(ctx, uri, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &DataBillDownloadUrlQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单线下交易预创建 alipay.trade.precreate
// StatusCode = 200 is success
func (a *ClientV3) TradePrecreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradePrecreateRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradePrecreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradePrecreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradePrecreateRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 统一收单交易创建接口 alipay.trade.create
// StatusCode = 200 is success
func (a *ClientV3) TradeCreate(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeCreateRsp, err error) {
	err = bm.CheckEmptyError("out_trade_no", "total_amount", "subject", "product_code", "op_app_id")
	if err != nil {
		return nil, err
	}
	if bm.GetString("buyer_id") == gopay.NULL && bm.GetString("buyer_open_id") == gopay.NULL {
		return nil, errors.New("buyer_id and buyer_open_id are not allowed to be null at the same time")
	}
	authorization, err := a.authorization(MethodPost, v3TradeCreate, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeCreate, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeCreateRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 支付宝订单信息同步接口 alipay.trade.orderinfo.sync
// StatusCode = 200 is success
func (a *ClientV3) TradeOrderInfoSync(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderInfoSyncRsp, err error) {
	err = bm.CheckEmptyError("trade_no", "out_request_no", "biz_type")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3TradeOrderInfoSync, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3TradeOrderInfoSync, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &TradeOrderInfoSyncRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 刷脸支付初始化 zoloz.authentication.smilepay.initialize
// StatusCode = 200 is success
func (a *ClientV3) ZolozAuthenticationSmilepayInitialize(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationSmilepayInitializeRsp, err error) {
	authorization, err := a.authorization(MethodPost, v3ZolozAuthenticationSmilepayInitialize, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZolozAuthenticationSmilepayInitialize, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZolozAuthenticationSmilepayInitializeRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}

// 查询刷脸结果信息接口 zoloz.authentication.customer.ftoken.query
// StatusCode = 200 is success
func (a *ClientV3) ZolozAuthenticationCustomerFtokenQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ZolozAuthenticationCustomerFtokenQueryRsp, err error) {
	err = bm.CheckEmptyError("ftoken", "biz_type")
	if err != nil {
		return nil, err
	}
	authorization, err := a.authorization(MethodPost, v3ZolozAuthenticationCustomerFtokenQuery, bm)
	if err != nil {
		return nil, err
	}
	res, bs, err := a.doPost(ctx, bm, v3ZolozAuthenticationCustomerFtokenQuery, authorization)
	if err != nil {
		return nil, err
	}
	aliRsp = &ZolozAuthenticationCustomerFtokenQueryRsp{StatusCode: res.StatusCode}
	if res.StatusCode != http.StatusOK {
		if err = json.Unmarshal(bs, &aliRsp.ErrResponse); err != nil {
			return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
		}
		return aliRsp, nil
	}
	if err = json.Unmarshal(bs, aliRsp); err != nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	return aliRsp, a.autoVerifySignByCert(res, bs)
}
