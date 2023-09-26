package wechat

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/go-pay/gopay"
)

// 公众号纯签约（正式）
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_1.shtml
func (w *Client) EntrustPublic(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustPublicResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(ctx, bm, entrustPublic, SignType_MD5)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPublicResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// APP纯签约-预签约接口-获取预签约ID（正式）
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_2.shtml
func (w *Client) EntrustAppPre(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustAppPreResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(ctx, bm, entrustApp)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustAppPreResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// H5纯签约（正式）
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_4.shtml
func (w *Client) EntrustH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustH5Response, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp", "clientip")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(ctx, bm, entrustH5, SignType_HMAC_SHA256)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustH5Response)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 支付中签约（正式）
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_5.shtml
func (w *Client) EntrustPaying(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustPayingResponse, err error) {
	err = bm.CheckEmptyError("contract_mchid", "contract_appid",
		"out_trade_no", "nonce_str", "body", "notify_url", "total_fee",
		"spbill_create_ip", "trade_type", "plan_id", "contract_code",
		"request_serial", "contract_display_account", "contract_notify_url")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(ctx, bm, entrustPaying)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPayingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 申请扣款
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_8.shtml
func (w *Client) EntrustApplyPay(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustApplyPayResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "body", "out_trade_no", "total_fee", "notify_url", "trade_type", "contract_id")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(ctx, bm, entrustApplyPay)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustApplyPayResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 申请解约
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_9.shtml
func (w *Client) EntrustDelete(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustDeleteResponse, err error) {
	err = bm.CheckEmptyError("contract_termination_remark", "version")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(ctx, bm, entrustDelete)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustDeleteResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询签约关系
// 文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_7.shtml
func (w *Client) EntrustQuery(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustQueryResponse, err error) {
	err = bm.CheckEmptyError("version")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(ctx, bm, entrustQuery)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustQueryResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}
