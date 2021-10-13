package wechat

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/cedarwu/gopay"
	"github.com/cedarwu/gopay/pkg/util"
)

// 统一下单
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_1.shtml
func (w *Client) UnifiedOrder(bm gopay.BodyMap) (wxRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "notify_url", "trade_type")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, unifiedOrder, nil)
	} else {
		bm.Set("total_fee", 101)
		bs, err = w.doSanBoxPost(bm, sandboxUnifiedOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(UnifiedOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 提交付款码支付
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter4_1.shtml
func (w *Client) Micropay(bm gopay.BodyMap) (wxRsp *MicropayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "spbill_create_ip", "auth_code")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, microPay, nil)
	} else {
		bm.Set("total_fee", 1)
		bs, err = w.doSanBoxPost(bm, sandboxMicroPay)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(MicropayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 查询订单
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_2.shtml
func (w *Client) QueryOrder(bm gopay.BodyMap) (wxRsp *QueryOrderResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("transaction_id") == util.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, orderQuery, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxOrderQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 关闭订单
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_3.shtml
func (w *Client) CloseOrder(bm gopay.BodyMap) (wxRsp *CloseOrderResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, closeOrder, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxCloseOrder)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(CloseOrderResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 申请退款
//	注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_4.shtml
func (w *Client) Refund(bm gopay.BodyMap) (wxRsp *RefundResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "total_fee", "refund_fee")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("out_trade_no") == util.NULL && bm.GetString("transaction_id") == util.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(nil, nil, nil); err != nil {
			return nil, nil, err
		}
		bs, err = w.doProdPost(bm, refund, tlsConfig)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxRefund)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(RefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 查询退款
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_5.shtml
func (w *Client) QueryRefund(bm gopay.BodyMap) (wxRsp *QueryRefundResponse, resBm gopay.BodyMap, err error) {
	err = bm.CheckEmptyError("nonce_str")
	if err != nil {
		return nil, nil, err
	}
	if bm.GetString("refund_id") == util.NULL && bm.GetString("out_refund_no") == util.NULL && bm.GetString("transaction_id") == util.NULL && bm.GetString("out_trade_no") == util.NULL {
		return nil, nil, errors.New("refund_id, out_refund_no, out_trade_no, transaction_id are not allowed to be null at the same time")
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(bm, refundQuery, nil)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxRefundQuery)
	}
	if err != nil {
		return nil, nil, err
	}
	wxRsp = new(QueryRefundResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalStruct(%s)：%w", string(bs), err)
	}
	resBm = make(gopay.BodyMap)
	if err = xml.Unmarshal(bs, &resBm); err != nil {
		return nil, nil, fmt.Errorf("xml.UnmarshalBodyMap(%s)：%w", string(bs), err)
	}
	return wxRsp, resBm, nil
}

// 撤销订单
//	注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter4_3.shtml
func (w *Client) Reverse(bm gopay.BodyMap) (wxRsp *ReverseResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(nil, nil, nil); err != nil {
			return nil, err
		}
		bs, err = w.doProdPost(bm, reverse, tlsConfig)
	} else {
		bs, err = w.doSanBoxPost(bm, sandboxReverse)
	}
	if err != nil {
		return nil, err
	}
	wxRsp = new(ReverseResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
