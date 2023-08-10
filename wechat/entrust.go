package wechat

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/cedarwu/gopay"
)

// 公众号纯签约（正式）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_1.shtml
func (w *Client) EntrustPublic(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustPublicResponse, header http.Header, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, nil, err
	}
	bs, header, err := w.doProdGet(ctx, bm, entrustPublic, SignType_MD5)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(EntrustPublicResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}

// APP纯签约-预签约接口-获取预签约ID（正式）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_2.shtml
func (w *Client) EntrustAppPre(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustAppPreResponse, header http.Header, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp")
	if err != nil {
		return nil, nil, err
	}
	bs, _, _, header, err := w.doProdPost(ctx, bm, entrustApp, nil)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(EntrustAppPreResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}

// H5纯签约（正式）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_4.shtml
func (w *Client) EntrustH5(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustH5Response, header http.Header, err error) {
	err = bm.CheckEmptyError("plan_id", "contract_code", "request_serial", "contract_display_account", "notify_url", "version", "timestamp", "clientip")
	if err != nil {
		return nil, nil, err
	}
	bs, header, err := w.doProdGet(ctx, bm, entrustH5, SignType_HMAC_SHA256)
	if err != nil {
		return nil, header, err
	}
	wxRsp = new(EntrustH5Response)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, header, nil
}

// 支付中签约（正式）
//
//	文档地址：https://pay.weixin.qq.com/wiki/doc/api/wxpay_v2/papay/chapter3_5.shtml
func (w *Client) EntrustPaying(ctx context.Context, bm gopay.BodyMap) (wxRsp *EntrustPayingResponse, bs []byte, url string, statusCode int, header http.Header, err error) {
	err = bm.CheckEmptyError("contract_mchid", "contract_appid",
		"out_trade_no", "nonce_str", "body", "notify_url", "total_fee",
		"spbill_create_ip", "trade_type", "plan_id", "contract_code",
		"request_serial", "contract_display_account", "contract_notify_url")
	if err != nil {
		return nil, nil, "", 0, nil, err
	}
	bs, url, statusCode, header, err = w.doProdPost(ctx, bm, entrustPaying, nil)
	if err != nil {
		return nil, nil, "", 0, header, err
	}
	wxRsp = new(EntrustPayingResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, bs, url, statusCode, header, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, bs, url, statusCode, header, nil
}
