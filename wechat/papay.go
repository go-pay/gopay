package wechat

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/go-pay/gopay"
)

// 公众号纯签约（正式）
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011986768
func (w *Client) EntrustPublic(ctx context.Context, bm gopay.BodyMap) (entrustURL string, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return gopay.NULL, err
	}
	// https://api.mch.weixin.qq.com/papay/entrustweb?
	// appid=wx426a3015555a46be&contract_code=122&contract_display_account=name1&mch_id=1223816102¬ify_url=http%3A%2F%2Fwww.qq.com%2Ftest%2Fpapay&plan_id=106&request_serial=123×tamp=1414488825&version=1.0&sign=FF1A406564EE701064450CA2149E2514
	if bm.GetString("appid") == gopay.NULL {
		bm.Set("appid", w.AppId)
	}
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", w.MchId)
	}
	bm.Remove("sign")
	sign := w.getReleaseSign(w.ApiKey, SignType_MD5, bm)
	bm.Set("sign", sign)
	url := baseUrlCh + entrustPublic
	if w.BaseURL != gopay.NULL {
		url = w.BaseURL + entrustPublic
	}
	entrustURL = url + "?" + bm.EncodeURLParams()
	if w.DebugSwitch == gopay.DebugOn {
		w.logger.Debugf("EntrustPublic_URL: %s", entrustURL)
	}
	return entrustURL, nil
}

// APP纯签约-预签约接口-获取预签约ID（正式）
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011986804
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
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011987295
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
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011987320
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
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011987377
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
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011987432
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
// 文档地址：https://pay.weixin.qq.com/doc/v2/merchant/4011987640
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
