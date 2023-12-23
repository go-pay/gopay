/*
 微信现金红包
 文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_1
*/

package wechat

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/go-pay/gopay"
)

// 发放现金红包
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，total_num = 1
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (w *Client) SendCashRed(ctx context.Context, bm gopay.BodyMap) (wxRsp *SendCashRedResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "wishing", "client_ip", "act_name", "remark")
	if err != nil {
		return nil, err
	}
	if bm.GetString("wxappid") == gopay.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := w.getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	bs, err := w.doProdPostPureTLS(ctx, bm, sendCashRed)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 发放现金裂变红包
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，amt_type = ALL_RAND
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (w *Client) SendGroupCashRed(ctx context.Context, bm gopay.BodyMap) (wxRsp *SendCashRedResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "amt_type", "wishing", "act_name", "remark")
	if err != nil {
		return nil, err
	}

	if bm.GetString("wxappid") == gopay.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := w.getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	bs, err := w.doProdPostPureTLS(ctx, bm, sendGroupCashRed)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 发放小程序红包
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，total_num = 1，notify_way = MINI_PROGRAM_JSAPI
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=3
func (w *Client) SendAppletRed(ctx context.Context, bm gopay.BodyMap) (wxRsp *SendAppletRedResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "wishing", "act_name", "remark", "notify_way")
	if err != nil {
		return nil, err
	}

	if bm.GetString("wxappid") == gopay.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := w.getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	bs, err := w.doProdPostPureTLS(ctx, bm, sendAppletRed)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendAppletRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}

// 查询红包记录
// 注意：请在初始化client时，调用 client 添加证书的相关方法添加证书
// 注意：此处参数中的 appid 需要单独传参，不复用 NewClient 时的 appid，bill_type = MCHT
// 微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
func (w *Client) QueryRedRecord(ctx context.Context, bm gopay.BodyMap) (wxRsp *QueryRedRecordResponse, err error) {
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "appid", "bill_type")
	if err != nil {
		return nil, err
	}

	if bm.GetString("appid") == gopay.NULL {
		bm.Set("appid", w.AppId)
	}
	if bm.GetString("mch_id") == gopay.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == gopay.NULL {
		sign := w.getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	bs, err := w.doProdPostPureTLS(ctx, bm, getRedRecord)
	if err != nil {
		return nil, err
	}
	wxRsp = new(QueryRedRecordResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", gopay.UnmarshalErr, err, string(bs))
	}
	return wxRsp, nil
}
