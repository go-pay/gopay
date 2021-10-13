package wechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cedarwu/gopay"
)

// 查询特约商户账户实时余额API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter5_1.shtml
func (c *ClientV3) V3EcommerceBalance(subMchid string) (*EcommerceBalanceRsp, error) {
	url := fmt.Sprintf(v3EcommerceBalance, subMchid)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &EcommerceBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(EcommerceBalance)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询账户实时余额API
//	Code = 0 is success
// 	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter5_1.shtml
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter5_2.shtml
func (c *ClientV3) V3MerchantBalance(accountType string) (*MerchantBalanceRsp, error) {
	url := fmt.Sprintf(v3MerchantBalance, accountType)
	authorization, err := c.authorization(MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(url, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantBalance)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 查询账户日终余额API
//	date示例值：2019-08-17
//	Code = 0 is success
// 	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter5_2.shtml
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter5_3.shtml
func (c *ClientV3) V3MerchantDayBalance(accountType, date string) (*MerchantBalanceRsp, error) {
	uri := fmt.Sprintf(v3MerchantDayBalance, accountType) + "?date=" + date
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantBalanceRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantBalance)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 特约商户银行来账查询API
//	Code = 0 is success
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter3_6.shtml
func (c *ClientV3) V3EcommerceIncomeRecord(bm gopay.BodyMap) (*PartnerIncomeRecordRsp, error) {
	uri := v3EcommerceIncomeRecord + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &PartnerIncomeRecordRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(PartnerIncomeRecord)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}

// 商户/服务商银行来账查询API
//	Code = 0 is success
// 	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer/chapter3_7.shtml
// 	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3/wxpay/pay/transfer_partner/chapter3_7.shtml
func (c *ClientV3) V3MerchantIncomeRecord(bm gopay.BodyMap) (*MerchantIncomeRecordRsp, error) {
	uri := v3MerchantIncomeRecord + "?" + bm.EncodeURLParams()
	authorization, err := c.authorization(MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	res, si, bs, err := c.doProdGet(uri, authorization)
	if err != nil {
		return nil, err
	}

	wxRsp := &MerchantIncomeRecordRsp{Code: Success, SignInfo: si}
	wxRsp.Response = new(MerchantIncomeRecord)
	if err = json.Unmarshal(bs, wxRsp.Response); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s)：%w", string(bs), err)
	}
	if res.StatusCode != http.StatusOK {
		wxRsp.Code = res.StatusCode
		wxRsp.Error = string(bs)
		return wxRsp, nil
	}
	return wxRsp, c.verifySyncSign(si)
}
