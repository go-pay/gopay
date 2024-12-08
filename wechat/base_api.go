package wechat

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/go-pay/gopay"
)

// 统一下单
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_1
func (w *Client) UnifiedOrder(ctx context.Context, bm gopay.BodyMap) (wxRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, unifiedOrder)
	} else {
		bm.Set("total_fee", 101)
		bs, err = w.doSanBoxPost(ctx, bm, sandboxUnifiedOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 提交付款码支付
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_10&index=1
func (w *Client) Micropay(ctx context.Context, bm gopay.BodyMap) (wxRsp *MicropayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "auth_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, microPay)
	} else {
		bm.Set("total_fee", 1)
		bs, err = w.doSanBoxPost(ctx, bm, sandboxMicroPay)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(MicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询订单
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_2
func (w *Client) QueryOrder(ctx context.Context, bm gopay.BodyMap) (wxRsp *QueryOrderResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, orderQuery)
	} else {
		bs, err = w.doSanBoxPost(ctx, bm, sandboxOrderQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s): %w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s): %w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 关闭订单
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_3
func (w *Client) CloseOrder(ctx context.Context, bm gopay.BodyMap) (wxRsp *CloseOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, closeOrder)
	} else {
		bs, err = w.doSanBoxPost(ctx, bm, sandboxCloseOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 申请退款
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_4
func (w *Client) Refund(ctx context.Context, bm gopay.BodyMap) (wxRsp *RefundResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "total_fee", "refund_fee")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("out_trade_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var (
		bs []byte
	)
	if w.IsProd {
		bs, err = w.doProdPostTLS(ctx, bm, refund)
	} else {
		bs, err = w.doSanBoxPost(ctx, bm, sandboxRefund)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s): %w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s): %w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 查询退款
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/jsapi_sl.php?chapter=9_5
func (w *Client) QueryRefund(ctx context.Context, bm gopay.BodyMap) (wxRsp *QueryRefundResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("refund_id") == gopay.NULL && bm.GetString("out_refund_no") == gopay.NULL && bm.GetString("transaction_id") == gopay.NULL && bm.GetString("out_trade_no") == gopay.NULL {
		return nil, nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, refundQuery)
	} else {
		bs, err = w.doSanBoxPost(ctx, bm, sandboxRefundQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s): %w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s): %w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 撤销订单
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 商户文档：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
// 服务商文档：https://pay.weixin.qq.com/wiki/doc/api/micropay_sl.php?chapter=9_11&index=3
func (w *Client) Reverse(ctx context.Context, bm gopay.BodyMap) (wxRsp *ReverseResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var (
		bs []byte
	)
	if w.IsProd {
		bs, err = w.doProdPostTLS(ctx, bm, reverse)
	} else {
		bs, err = w.doSanBoxPost(ctx, bm, sandboxReverse)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}
