package wechat

import (
	"encoding/xml"
	"fmt"

	"github.com/iGoogle-ink/gopay"
)

// 公众号纯签约（正式）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_1.shtml
func (w *Client) EntrustPublic(bm gopay.BodyMap) (wxRsp *EntrustPublicResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(bm, entrustPublic, SignType_MD5)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPublicResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// APP纯签约-预签约接口-获取预签约ID（正式）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_2.shtml
func (w *Client) EntrustAppPre(bm gopay.BodyMap) (wxRsp *EntrustAppPreResponse, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, entrustApp, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustAppPreResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// H5纯签约（正式）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_4.shtml
func (w *Client) EntrustH5(bm gopay.BodyMap) (wxRsp *EntrustH5Response, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp", "clientip")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdGet(bm, entrustH5, SignType_HMAC_SHA256)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustH5Response)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// 支付中签约（正式）
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_5.shtml
func (w *Client) EntrustPaying(bm gopay.BodyMap) (wxRsp *EntrustPayingResponse, err error) {
	err = bm.CheckEmptyError("contract_mchid", "contract_appid",
		"out_trade_no", "nonce_str", "body", "notify_url", "total_fee",
		"spbill_create_ip", "trade_type", "plan_id", "contract_code",
		"request_serial", "contract_display_account", "contract_notify_url")
	if err != nil {
		return nil, err
	}
	bs, err := w.doProdPost(bm, entrustPaying, nil)
	if err != nil {
		return nil, err
	}
	wxRsp = new(EntrustPayingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
