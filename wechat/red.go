/*
	微信现金红包
	文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_1
*/

package wechat

import (
	"encoding/xml"
	"fmt"

	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/pkg/util"
)

// SendCashRed 发放现金红包
//	注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，total_num = 1
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_4&index=3
func (w *Client) SendCashRed(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *SendCashRedResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "wishing", "client_ip", "act_name", "remark")
	if err != nil {
		return nil, err
	}
	if bm.GetString("wxappid") == util.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == util.NULL {
		sign := getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPostPure(bm, sendCashRed, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// SendGroupCashRed 发放现金裂变红包
//	注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，amt_type = ALL_RAND
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_5&index=4
func (w *Client) SendGroupCashRed(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *SendCashRedResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "amt_type", "wishing", "act_name", "remark")
	if err != nil {
		return nil, err
	}

	if bm.GetString("wxappid") == util.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == util.NULL {
		sign := getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPostPure(bm, sendGroupCashRed, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendCashRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// SendAppletRed 发放小程序红包
//	注意：此处参数中的 wxappid 需要单独传参，不复用 NewClient 时的 appid，total_num = 1，notify_way = MINI_PROGRAM_JSAPI
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=3
func (w *Client) SendAppletRed(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *SendAppletRedResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "wxappid", "send_name", "re_openid", "total_amount", "total_num", "wishing", "act_name", "remark", "notify_way")
	if err != nil {
		return nil, err
	}

	if bm.GetString("wxappid") == util.NULL {
		bm.Set("wxappid", w.AppId)
	}
	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == util.NULL {
		sign := getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPostPure(bm, sendAppletRed, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(SendAppletRedResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}

// QueryRedRecord 查询红包记录
//	注意：此处参数中的 appid 需要单独传参，不复用 NewClient 时的 appid，bill_type = MCHT
//	微信文档：https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=13_6&index=5
func (w *Client) QueryRedRecord(bm gopay.BodyMap, certFilePath, keyFilePath, pkcs12FilePath interface{}) (wxRsp *QueryRedRecordResponse, err error) {
	if err = checkCertFilePathOrContent(certFilePath, keyFilePath, pkcs12FilePath); err != nil {
		return nil, err
	}
	err = bm.CheckEmptyError("nonce_str", "mch_billno", "appid", "bill_type")
	if err != nil {
		return nil, err
	}

	if bm.GetString("appid") == util.NULL {
		bm.Set("appid", w.AppId)
	}
	if bm.GetString("mch_id") == util.NULL {
		bm.Set("mch_id", w.MchId)
	}
	if bm.GetString("sign") == util.NULL {
		sign := getReleaseSign(w.ApiKey, SignType_MD5, bm)
		bm.Set("sign", sign)
	}

	tlsConfig, err := w.addCertConfig(certFilePath, keyFilePath, pkcs12FilePath)
	if err != nil {
		return nil, err
	}

	bs, err := w.doProdPostPure(bm, getRedRecord, tlsConfig)
	if err != nil {
		return nil, err
	}
	wxRsp = new(QueryRedRecordResponse)
	if err = xml.Unmarshal(bs, wxRsp); err != nil {
		return nil, fmt.Errorf("xml.Unmarshal(%s)：%w", string(bs), err)
	}
	return wxRsp, nil
}
