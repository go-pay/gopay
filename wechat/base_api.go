package wechat

import (
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gotil"
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
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL {
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
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter3_4.shtml
func (w *Client) Refund(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *RefundResponse, resBm gopay.BodyMap, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_refund_no", "total_fee", "refund_fee")
	if err != nil {
		return nil, nil, err
	}
	if bm.Get("out_trade_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL {
		return nil, nil, errors.New("out_trade_no and transaction_id are not allowed to be null at the same time")
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
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
	if bm.Get("refund_id") == gotil.NULL && bm.Get("out_refund_no") == gotil.NULL && bm.Get("transaction_id") == gotil.NULL && bm.Get("out_trade_no") == gotil.NULL {
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
//	注意：如已使用client.AddCertFilePath()或client.AddCertFileContent()添加过证书，参数certFilePath、keyFilePath、pkcs12FilePath全传 nil，否则，3证书Path均不可空
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/open/chapter4_3.shtml
func (w *Client) Reverse(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *ReverseResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "out_trade_no")
	if err != nil {
		return nil, err
	}
	var (
		bs        []byte
		tlsConfig *tls.Config
	)
	if w.IsProd {
		if tlsConfig, err = w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
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
